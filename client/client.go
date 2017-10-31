package client

import (
	"log"
	"time"

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

func Echo(message string) string {
	url := nats.DefaultURL
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}
	defer nc.Close()

	subj := "echo"
	payload := []byte(message)

	msg, err := nc.Request(subj, []byte(payload), 100*time.Millisecond)
	if err != nil {
		if nc.LastError() != nil {
			log.Fatalf("Error in Request: %v\n", nc.LastError())
		}
		log.Fatalf("Error in Request: %v\n", err)
	}

	log.Printf("Published [%s] : '%s'\n", subj, payload)
	log.Printf("Received [%v] : '%s'\n", msg.Subject, string(msg.Data))

	return string(msg.Data)
}
