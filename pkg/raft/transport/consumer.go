package transport

import "github.com/luxordynamics/luxor/pkg/raft/protocol/out"

// We could apply a nasty hack to only have one consumer for all packets.
// To do this we would have to create a Packet interface which only defines the method GetID(uint8)
// The next step is to define the id attribute in our protobuf files.
// Because the protobuf compiler will generate a GetID method for the field, all the conditions for
// the protobuf struct to implement the Packet interface would be satisfied.
// Now we can create a single consumer like this:
//
// type Consumer interface {
//     Consume(chan<- Packet)
//}

type LogConsumer interface {

	// ConsumeResponse consumes all AppendEntriesResponse packets.
	ConsumeResponse(chan<- raft.AppendEntriesResponse)

	// ConsumeRequest consumes all AppendEntriesRequest packets.
	ConsumeRequest(chan<- raft.AppendEntriesRequest)

	// Closes the connection and frees all resources.
	Close()
}

type VoteConsumer interface {

	// ConsumeVoteResponse consumes all VoteRequest packets.
	ConsumeVoteResponse(<-chan raft.VoteResponse)

	// ConsumeVoteRequest consumes all VoteRequest packets.
	ConsumeVoteRequest(<-chan raft.VoteRequest)

	// Closes the connection and frees all resources.
	Close()
}
