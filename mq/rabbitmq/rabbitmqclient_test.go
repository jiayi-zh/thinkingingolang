package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

func Test_RabbitMQConsumer(t *testing.T) {
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

	err = rabbitMqClient.RegisterConsumer(&Consumer{
		Queue:       "queue_test",
		ConsumerTag: "测试消费者",
		ConsumeFun: func(delivery *amqp.Delivery) {
			fmt.Println("receive: ", string(delivery.Body))
			time.Sleep(time.Second * 3)
			err = delivery.Ack(false)
			if err != nil {
				fmt.Println("ack fail, cause: ", fmt.Sprintf("%v", err))
			}
		},
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("######## register consumer fail, cause: %v", err))
	}

	select {}
}

func Test_RabbitMQPublish(t *testing.T) {
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

	obj := &Obj{
		Id:   1,
		Name: "zy",
		Age:  25,
		Sex:  0,
		City: "甘肃",
	}

	for now := range time.Tick(time.Millisecond) {
		obj.TimeStamp = now.Unix()
		err = rabbitMqClient.PublishJson("exchange_test", "key_test", obj)
		if err != nil {
			fmt.Println(fmt.Sprintf("######## publish message fail, cause: %v", err))
		}
	}
}

type Obj struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Sex       int    `json:"sex"`
	City      string `json:"city"`
	TimeStamp int64  `json:"timeStamp"`
}
