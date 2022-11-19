package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	go subToPewdiepie()

	for {

	}
}

func subToPewdiepie() {
	// Connect to a server
	nc, _ := nats.Connect("nats://95.165.107.100:4222")
	defer nc.Drain()

	// Simple Publisher
	// nc.Publish("foo", []byte("Hello World"))

	// Responding to a request message
	nc.Subscribe("ith", func(m *nats.Msg) {
		m.Respond([]byte("Фамилия принял!"))
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	nc.Publish("ith", []byte("Hello there bruv"))

	// Close connection
	// nc.Close()
}
