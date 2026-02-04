package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func consumerMain() {
	fmt.Println("Consumer Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Conected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	msg, err := ch.Consume(
		"TestQueue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Printf("Recived Msg: %s\n", d.Body)
		}
	}()

	fmt.Println("Waiting for the message...")
	<-forever
}
