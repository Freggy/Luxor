package transport

import (
	"log"
	"testing"
)

func TestAMQPTransport(t *testing.T) {
	trans := NewAMQPTransport(Config{
		Host: "ramses.freggy.de",
		Port: 5672,
		Password: "guest",
		User: "guest",
		UseTLS: false,
		ExtraData: map[string]string {
			"queue_id": "MeineIDYAAAAAAAAAAAAAA",
		},
	})

	err := trans.Connect()
	log.Println(err)
}
