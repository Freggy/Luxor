package transport

import "github.com/luxordynamics/luxor/pkg/raft/protocol/gen"

type Producer interface {

	// Sends the given packet over the wire
	Produce(<-chan raft.Packet)

	// Closes the connection and frees all resources.
	Close()
}
