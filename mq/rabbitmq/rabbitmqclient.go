package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"strings"
	"time"
)

const (
	RabbitMQDefaultHost        = "127.0.0.1"
	RabbitMQDefaultUsername    = "guest"
	RabbitMQDefaultPassword    = "guest"
	RabbitMQDefaultPort        = 5672
	RabbitMQDefaultVirtualHost = "/"
	RabbitMQDefaultRoutingKey  = "#"
)

type RabbitmqClient struct {
	channel       *amqp.Channel
	cacheExchange []*Exchange
	cacheQueue    []*Queue
	cacheBinding  []*Binding
	cacheConsumer []*Consumer
}

const (
	Direct  = "direct"
	Topic   = "topic"
	Fanout  = "fanout"
	Headers = "headers"
)

func (rc *RabbitmqClient) Init() {
	rc.cacheExchange = make([]*Exchange, 0, 0)
	rc.cacheQueue = make([]*Queue, 0, 0)
	rc.cacheBinding = make([]*Binding, 0, 0)
	rc.cacheConsumer = make([]*Consumer, 0, 0)
}

func (rc *RabbitmqClient) IsConnect() bool {
	return rc.channel != nil
}

func (rc *RabbitmqClient) Connect(rmd *RabbitMQMetadata) error {
	if rmd == nil {
		return errors.New("rabbitmq connect metadata must not blank")
	}
	rmd.verityAndFillDefault()
	go rc.retryConnectRabbitmq(rmd)
	return nil
}

func (rc *RabbitmqClient) retryConnectRabbitmq(rmd *RabbitMQMetadata) {
	uri := rmd.buildConnURI()
	for range time.Tick(3 * time.Second) {
		// build tcp connect
		conn, err := amqp.Dial(uri)
		if err != nil {
			fmt.Println(fmt.Sprintf("build rabbitmq connection fail, cause: %v", err))
			continue
		}
		// recovery
		exceptionConn := conn.NotifyClose(make(chan *amqp.Error))
		go func() {
			for connErr := range exceptionConn {
				rc.channel = nil
				fmt.Println(fmt.Sprintf("rabbitmq exception connection disconnect, cause: %v", connErr))
				go rc.retryConnectRabbitmq(rmd)
			}
		}()

		// build channel connect
		channel, err := conn.Channel()
		if err != nil {
			fmt.Println(fmt.Sprintf("build rabbitmq channel fail, cause: %v", err))
			continue
		}
		// recovery
		exceptionChn := channel.NotifyClose(make(chan *amqp.Error))
		go func() {
			for connErr := range exceptionChn {
				rc.channel = nil
				fmt.Println(fmt.Sprintf("rabbitmq exception channel disconnect, cause: %v", connErr))
				go rc.retryConnectRabbitmq(rmd)
			}
		}()
		rc.channel = channel

		fmt.Println("rabbitmq connect success")
		// reload metadata
		rc.playbackCacheMetadata()
		break
	}
}

func (rc *RabbitmqClient) DeclareExchange(exchange *Exchange) error {
	if exchange == nil || !exchange.verityAndFillDefault() {
		return errors.New("invalid parameter")
	}

	// cache
	rc.cacheExchange = append(rc.cacheExchange, exchange)

	if rc.channel == nil {
		return nil
	}

	return rc.channel.ExchangeDeclare(exchange.Name, exchange.Type, exchange.Durable, exchange.AutoDelete, exchange.Internal, exchange.NoWait, exchange.Args)
}

func (rc *RabbitmqClient) DeclareQueue(queue *Queue) error {
	if queue == nil || !queue.verityAndFillDefault() {
		return errors.New("invalid parameter")
	}

	// cache
	rc.cacheQueue = append(rc.cacheQueue, queue)

	if rc.channel == nil {
		return nil
	}

	_, err := rc.channel.QueueDeclare(queue.Name, queue.Durable, queue.AutoDelete, queue.Exclusive, queue.NoWait, queue.Args)
	return err
}

func (rc *RabbitmqClient) DeclareBinding(binding *Binding) error {
	if binding == nil {
		return errors.New("invalid parameter")
	}

	// cache
	rc.cacheBinding = append(rc.cacheBinding, binding)

	if rc.channel == nil {
		return nil
	}

	return rc.channel.QueueBind(binding.Queue, binding.RoutingKey, binding.Exchange, binding.NoWait, binding.Args)
}

func (rc *RabbitmqClient) RegisterConsumer(consumer *Consumer) error {
	if len(strings.TrimSpace(consumer.Queue)) == 0 || consumer.ConsumeFun == nil {
		return errors.New("invalid parameter")
	}

	// cache
	rc.cacheConsumer = append(rc.cacheConsumer, consumer)

	if rc.channel == nil {
		return nil
	}

	deliveries, err := rc.channel.Consume(consumer.Queue, consumer.ConsumerTag, consumer.AutoAck, consumer.Exclusive, consumer.NoLocal, consumer.NoWait, consumer.Args)
	if err != nil {
		return err
	}
	for {
		select {
		case delivery := <-deliveries:
			go consumer.ConsumeFun(&delivery)
		}
	}
}

func (rc *RabbitmqClient) Publish2(exchange, key string, mandatory, immediate bool, publish amqp.Publishing) error {
	if rc.channel == nil {
		return errors.New("the connection task is not complete")
	}
	if err := rc.channel.Publish(exchange, key, mandatory, immediate, publish); err != nil {
		return err
	}
	return nil
}

func (rc *RabbitmqClient) PublishJson(exchange string, key string, obj interface{}) error {
	if obj == nil {
		return errors.New("invalid parameter")
	}
	if rc.channel == nil {
		return errors.New("the connection task is not complete")
	}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = rc.channel.Publish(exchange, key, false, false, amqp.Publishing{
		Timestamp:   time.Now(),
		ContentType: "application/json",
		Body:        jsonBytes,
	})
	return err
}

func (rc *RabbitmqClient) playbackCacheMetadata() {
	if rc.cacheExchange != nil && len(rc.cacheExchange) > 0 {
		for _, exchange := range rc.cacheExchange {
			err := rc.DeclareExchange(exchange)
			if err != nil {
				fmt.Println(fmt.Sprintf("declare Exchange fail, cause: %v", err))
			}
		}
	}
	if rc.cacheQueue != nil && len(rc.cacheQueue) > 0 {
		for _, queue := range rc.cacheQueue {
			err := rc.DeclareQueue(queue)
			if err != nil {
				fmt.Println(fmt.Sprintf("declare Queue fail, cause: %v", err))
			}
		}
	}
	if rc.cacheBinding != nil && len(rc.cacheBinding) > 0 {
		for _, binding := range rc.cacheBinding {
			err := rc.DeclareBinding(binding)
			if err != nil {
				fmt.Println(fmt.Sprintf("declare consumer fail, cause: %v", err))
			}
		}
	}
	if rc.cacheConsumer != nil && len(rc.cacheConsumer) > 0 {
		for _, consumer := range rc.cacheConsumer {
			err := rc.RegisterConsumer(consumer)
			if err != nil {
				fmt.Println(fmt.Sprintf("register consumer fail, cause: %v", err))
			}
		}
	}
}
