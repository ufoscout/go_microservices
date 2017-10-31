package main

import (
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {

	url := nats.DefaultURL
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	i := 0

	nc.Subscribe("echo", func(msg *nats.Msg) {
		i++
		printMsg(msg, i)
	})

	nc.Subscribe("hello", func(msg *nats.Msg) {
		i++
		printMsg(msg, i)
	})

	nc.Flush()

	log.Printf("Connected to Nats at '%s'\n", url)
	runtime.Goexit()
}
