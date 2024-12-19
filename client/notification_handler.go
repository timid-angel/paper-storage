package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func subscribeToNotification(address, exchangeName string, has_add *bool) {
	conn, err := amqp.Dial(address)
	if err != nil {
		log.Fatalln("failed to connect to rabbitmq service: " + err.Error())
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open a channel: " + err.Error())
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to declare exchange: " + err.Error())
	}

	queue, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatalln("failed to declare a queue: " + err.Error())
	}

	err = ch.QueueBind(queue.Name, "", exchangeName, false, nil)
	if err != nil {
		log.Fatalln("failed to bind queue to exchange: " + err.Error())
	}

	messageChannel, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("failed to register consumer: " + err.Error())
	}

	for msg := range messageChannel {
		if !*has_add {
			fmt.Printf("\n\n\t\033[90m> [NOTIFICATION] %s\n\n\033[0m", msg.Body)
			fmt.Printf("\033[0m\033[3;36m>\033[0m \033[1;3;35m")
		} else {
			*has_add = true
		}
	}
}
