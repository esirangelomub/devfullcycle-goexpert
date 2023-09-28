package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	// e.g. rabbitmq
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			//time.Sleep(1 * time.Second)
			msg := Message{i, "Hello from rabbitmq"}
			c1 <- msg
		}
	}()

	// e.g. kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			//time.Sleep(1 * time.Second)
			msg := Message{ID: i, Msg: "Hello from kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1: // e.g. rabbitmq
			fmt.Printf("Received from RabbitMQ: ID: %d - %+v\n", msg.ID, msg.Msg)

		case msg := <-c2: // e.g. kafka
			fmt.Printf("Received from Kafka: ID: %d - %+v\n", msg.ID, msg.Msg)

		case <-time.After(3 * time.Second):
			println("timeout")

			//default:
			//	println("default")
		}
	}
}
