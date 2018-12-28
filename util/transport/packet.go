package transport

import (
	"github.com/golang/protobuf/proto"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/gen"
)

type PacketContainer struct {

	// UUID of the client to reply to
	ReplyTo string

	// Unix timestamp in milliseconds when the packet was received
	Received int64

	// The actual payload
	Packet *raft.Packet
}

// NewAppendEntriesRequest creates a new AppendEntriesRequest (0x1) packet.
//
// term         - Leader’s term
// prevLogIndex - Index of log entry immediately preceding new ones
// leaderCommit - Term of prevLogIndex entry
// leaderId     - Id of the leader
// entries      - Log entries to store (empty for heartbeat; may send more than one for efficiency)
func NewAppendEntriesRequest(
	term uint32,
	prevLogIndex uint32,
	leaderCommit uint32,
	leaderId string,
	entries []*raft.Entry) *raft.Packet {

	return toPacket(0x1, &raft.AppendEntriesRequest{
		Term:         term,
		PrevLogIndex: prevLogIndex,
		LeaderCommit: leaderCommit,
		LeaderId:     leaderId,
		Entries:      entries,
	})
}

// NewAppendEntriesResponse creates a new AppendEntriesResponse (0x2) packet.
//
// term       - Followers current term.
// followerId - Follower that sent this response
// success    - True if follower contained entry matching prevLogIndex and prevLogTerm
func NewAppendEntriesResponse(
	term uint32,
	followerId string,
	success bool) *raft.Packet {
	return toPacket(0x2, &raft.AppendEntriesResponse{
		Term:       term,
		FollowerId: followerId,
		Success:    success,
	})
}

// NewVoteRequest creates a new VoteRequest (0x3) packet.
//
// term         - Candidates’s term.
// lastLogIndex - Candidate requesting vote.
// lastLogTerm  - Index of candidate's last log entry.
// candidateId  - Term of candidate's last log entry.
func NewVoteRequest(
	term uint32,
	lastLogIndex uint32,
	lastLogTerm uint32,
	candidateId string) *raft.Packet {

	return toPacket(0x3, &raft.VoteRequest{
		Term:         term,
		LastLogIndex: lastLogIndex,
		LastLogTerm:  lastLogTerm,
		CandidateId:  candidateId,
	})
}

// NewVoteResponse creates a new VoteResponse (0x4) packet.
//
// term        - Term of candidate.
// voteGranted - True means that candidate received vote.
func NewVoteResponse(term uint32, voteGranted bool) *raft.Packet {
	return toPacket(0x4, &raft.VoteResponse{
		Term:        term,
		VoteGranted: voteGranted,
	})
}

// toPacket creates the wrapper for the raft packet given.
//
// id  - Id of the packet
// msg - Protobuf message
func toPacket(id uint32, msg proto.Message) *raft.Packet {
	bytes, _ := proto.Marshal(msg)
	return &raft.Packet{
		Id:      id,
		Payload: bytes,
	}
}
