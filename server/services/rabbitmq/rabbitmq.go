package rabbitmq

import (
	"paper-server/server/domain"

	"github.com/streadway/amqp"
)

type RabbitmqService struct {
	address      string
	exchangeName string
}

func NewRabbitMqService(address string, exchangeName string) *RabbitmqService {
	return &RabbitmqService{
		address:      address,
		exchangeName: exchangeName,
	}
}

func (rmq *RabbitmqService) PublishNotification(message string) domain.IDomainError {
	conn, err := amqp.Dial(rmq.address)
	if err != nil {
		return domain.NewDomainError("failed to connect to RabbitMQ: " + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		domain.NewDomainError("failed to open channel: " + err.Error())
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		rmq.exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		domain.NewDomainError("failed to declare exchange: " + err.Error())
	}

	err = ch.Publish(rmq.exchangeName, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		domain.NewDomainError("failed to publish a message: " + err.Error())
	}

	return nil
}
