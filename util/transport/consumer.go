package transport

import (
	"github.com/golang/protobuf/proto"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/gen"
	"github.com/streadway/amqp"
	"log"
)

type Consumer interface {

	// Consume starts listening for packets.
	Consume() (chan<- raft.Packet, error)

	// Close closes the connection and frees all resources.
	Close() error
}

type AMQPConsumer struct {
	Queue   string
	Channel *amqp.Channel
}

func NewAMPQConsumer() *AMQPConsumer {
	return nil
}

func (c *AMQPConsumer) Consume() (chan<- raft.Packet, error) {
	pkts, err := c.Channel.Consume(
		c.Queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	pktChan := make(chan<- raft.Packet)

	go func() {
		for pd := range pkts {
			var p raft.Packet
			if err := proto.Unmarshal(pd.Body, &p); err != nil {
				log.Println(err)
				continue
			}
			pktChan <- p
		}
	}()
	return pktChan, nil
}

func (c *AMQPConsumer) Close() error {
	return c.Channel.Close()
}
