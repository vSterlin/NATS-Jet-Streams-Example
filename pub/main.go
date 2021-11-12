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
	js, err := nc.JetStream()

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

	_, err = js.Publish(DEMO+".test", []byte("Some message"))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer nc.Close()

}
