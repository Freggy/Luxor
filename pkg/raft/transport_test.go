package raft

import (
	"log"
	"testing"
)

func TestAMQPTransport(t *testing.T) {
	c := make(chan bool)

	trans1 := NewAMQPTransport(Config{
		Host: "ramses.freggy.de",
		Port: 5672,
		Password: "guest",
		User: "guest",
		UseTLS: false,
		ExtraData: map[string]string {
			"queue_id": "1",
		},
	})

	err := trans1.Connect()
	log.Println(err)

	err = trans1.StartConsume()
	log.Println(err)

	go func() {
		for d := range trans1.ConsumeChannel {
			log.Println(d.Packet)
		}
	}()

	trans2 := NewAMQPTransport(Config{
		Host: "ramses.freggy.de",
		Port: 5672,
		Password: "guest",
		User: "guest",
		UseTLS: false,
		ExtraData: map[string]string {
			"queue_id": "2",
		},
	})

	err = trans2.Connect()
	log.Println(err)

	err = trans2.StartProduce()
	log.Println(err)
	trans2.ProduceChannel <- &PacketContainer{
		"1",
		0,
		NewAppendEntriesResponse(1, "lol", false),
	}
	<- c
}
