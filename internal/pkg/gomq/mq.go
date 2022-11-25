package gomq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const MQURL = "amqp://test:123456@192.168.33.204:5672/"

type RabbitMQ struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
	Exchange   string
	RoutingKey string
	Mqurl      string
}

func NewRabbitMQ(queueName, exchange, routingKey string) *RabbitMQ {
	rabbitMQ := RabbitMQ{
		QueueName:  queueName,
		Exchange:   exchange,
		RoutingKey: routingKey,
		Mqurl:      MQURL,
	}
	var err error
	//创建rabbitmq连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	if err != nil {
		log.Fatalf("创建rabbitmq连接失败:%s", err)
	}
	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	if err != nil {
		log.Fatalf("创建Channel失败:%s", err)
	}
	return &rabbitMQ
}
