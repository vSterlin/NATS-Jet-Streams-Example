package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

const DEMO = "DEMO"

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	js, _ := nc.JetStream()

	stream, err := js.StreamInfo(DEMO)
	if err != nil {
		fmt.Println(err.Error())
	}

	if stream == nil {
		fmt.Println("Creating new stream")
		if _, err := js.AddStream(&nats.StreamConfig{
			Name:     DEMO,
			Subjects: []string{DEMO + ".*"},
		}); err != nil {
			fmt.Println(err)
		}
	}

	js.Subscribe(DEMO+".test", func(msg *nats.Msg) {
		fmt.Printf("msg.Data: %v\n", string(msg.Data))
		fmt.Printf("msg.Subject: %v\n", msg.Subject)
	})

	defer nc.Close()

	for {

	}
}
