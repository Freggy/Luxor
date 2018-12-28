package transport

import (
	"github.com/golang/protobuf/proto"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/gen"
)

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

func NewVoteResponse(term uint32, voteGranted bool) *raft.Packet {
	return toPacket(0x4, &raft.VoteResponse{
		Term:        term,
		VoteGranted: voteGranted,
	})
}

func toPacket(id uint32, msg proto.Message) *raft.Packet {
	bytes, _ := proto.Marshal(msg)
	return &raft.Packet{
		Id:      id,
		Payload: bytes,
	}
}
