package consumer

import (
	"fmt"
	"kafka/internal/shared"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	topic    string
	msgCH    chan<- string
}

func NewConsumer(msgCH chan<- string) *KafkaConsumer {

	cfg := shared.NewKafkaConfig()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Host,
		"group.id":          cfg.ConsumerGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{cfg.Topic}, nil)

	if err != nil {
		panic(err)
	}

	consumer := &KafkaConsumer{
		consumer: c,
		topic:    cfg.Topic,
		msgCH:    msgCH,
	}
	go consumer.readMsg()
	return consumer
}

func (c *KafkaConsumer) readMsg() {
	for {
		msg, err := c.consumer.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			// Send to channel for processing
			c.msgCH <- string(msg.Value)
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
	}
}

// Close closes the Kafka consumer connection
func (c *KafkaConsumer) Close() {
	c.consumer.Close()
}
