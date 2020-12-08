package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

type Obj struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func Test_RabbitMQClient(t *testing.T) {
	rabbitMqClient := new(RabbitmqClient)
	rabbitMqClient.Init()

	err := rabbitMqClient.Connect(&RabbitMQMetadata{
		Host:        "192.168.9.102",
		Port:        31005,
		Username:    "golang",
		Password:    "^JITV^8La4Np9BWD",
		VirtualHost: "server",
	})
	if err != nil {
		return
	}

	err = rabbitMqClient.DeclareExchange(&Exchange{
		Name: "exchange_test",
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("######## declare exchange fail, cause: %v", err))
	}

	err = rabbitMqClient.DeclareQueue(&Queue{
		Name: "queue_test",
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("######## declare queue fail, cause: %v", err))
	}

	err = rabbitMqClient.DeclareBinding(&Binding{
		Exchange:   "exchange_test",
		Queue:      "queue_test",
		RoutingKey: "key_test",
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("######## declare binding fail, cause: %v", err))
	}

	err = rabbitMqClient.RegisterConsumer(&Consumer{
		Queue: "queue_test",
		ConsumeFun: func(delivery *amqp.Delivery) {
			fmt.Println("receive: ", string(delivery.Body))
			err = delivery.Ack(false)
			if err != nil {
				fmt.Println("ack fail, cause: ", fmt.Sprintf("%v", err))
			}
		},
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("######## register consumer fail, cause: %v", err))
	}

	for range time.Tick(time.Second) {
		err = rabbitMqClient.PublishJson("exchange_test", "key_test", false, false, &Obj{Id: 1, Name: "zy"})
		if err != nil {
			fmt.Println(fmt.Sprintf("######## publish message fail, cause: %v", err))
		}
	}
}
