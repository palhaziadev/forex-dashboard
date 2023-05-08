package rabbbitmq_utils

import (
	"log"

	"github.com/streadway/amqp"
)

type RMQProducer struct {
	Queue            string
	ConnectionString string
}

func (x RMQProducer) OnError(err error, msg string) {
	if err != nil {
		log.Printf("Error occurred while publishing message on '%s' queue. Error message: %s", x.Queue, msg)
	}
}

func (x RMQProducer) PublishMessage(contentType string, body []byte) error {
	log.Printf(x.ConnectionString)
	conn, err := amqp.Dial(x.ConnectionString)
	x.OnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	x.OnError(err, "Failed to open a channel")
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		x.Queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	x.OnError(err, "Failed to declare a queue")
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		})
	x.OnError(err, "Failed to publish a message")
	if err != nil {
		return err
	}
	return nil
}
