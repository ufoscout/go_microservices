package client

import (
	"log"

	"github.com/nats-io/go-nats"
)

func Hello(name string) {
	url := nats.DefaultURL
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}
	defer nc.Close()

	subj := "hello"
	msg := []byte(name)

	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}

func echo() {
	url := nats.DefaultURL
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}
	defer nc.Close()

	subj := "echo"
	msg := []byte("Francesco")

	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
