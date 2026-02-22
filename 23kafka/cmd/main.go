package main

import (
	"fmt"
	"kafka/internal/consumer"
	"kafka/internal/producer"
	"time"
)

type Server struct {
	producer *producer.KafkaProducer
	consumer *consumer.KafkaConsumer
	msgCH    chan string
}

func NewServer() *Server {
	msgCH := make(chan string, 64)
	return &Server{
		producer: producer.NewKafkaProducer(""),
		consumer: consumer.NewConsumer(msgCH),
		msgCH:    msgCH,
	}
}

func (s *Server) PublishMsg() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	id := 1
	for t := range ticker.C {
		msg := fmt.Sprintf("Kafka message MessageID: %d, timestamp: %s", id, t.Format("15:04:05"))
		s.producer.Produce(msg)
		id++
	}
}

func (s *Server) handleMsg(msg string) {
	//db operations
	fmt.Printf("Recived Message: %s\n", msg)
}

func main() {
	s := NewServer()
	go s.PublishMsg()
	for msg := range s.msgCH {
		go s.handleMsg(msg)
	}
}
