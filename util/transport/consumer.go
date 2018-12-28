package transport

import (
	"github.com/golang/protobuf/proto"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/gen"
	"github.com/streadway/amqp"
	"log"
)

// We could apply a nasty hack to only have one consumer for all packets.
// To do this we would have to create a Packet interface which only defines the method GetID(uint8)
// The next step is to define the id attribute in our protobuf files.
// Because the protobuf compiler will generate a GetID method for the field, all the conditions for
// the protobuf struct to implement the Packet interface would be satisfied.
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
