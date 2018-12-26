package transport

import "github.com/luxordynamics/luxor/pkg/raft/protocol/gen"

type Packet interface {
	GetId() uint32
}

func NewAppendEntriesRequest(
	term uint32,
	prevLogIndex uint32,
	leaderCommit uint32,
	leaderId string,
	entries []*raft.Entry) *raft.AppendEntriesRequest {
	return &raft.AppendEntriesRequest{
		Id:           0x1,
		Term:         term,
		PrevLogIndex: prevLogIndex,
		LeaderCommit: leaderCommit,
		LeaderId:     leaderId,
		Entries:      entries,
	}
}

func NewAppendEntriesResponse(
	term uint32,
	followerId string,
	success bool) *raft.AppendEntriesResponse {
	return &raft.AppendEntriesResponse{
		Id:         0x2,
		Term:       term,
		FollowerId: followerId,
		Success:    success,
	}
}

func NewVoteRequest(
	term uint32,
	lastLogIndex uint32,
	lastLogTerm uint32,
	candidateId string) *raft.VoteRequest {
	return &raft.VoteRequest{
		Id:           0x3,
		Term:         term,
		LastLogIndex: lastLogIndex,
		LastLogTerm:  lastLogTerm,
		CandidateId:  candidateId,
	}
}

func NewVoteResponse(term uint32, voteGranted bool) *raft.VoteResponse {
	return &raft.VoteResponse{
		Id:          0x4,
		Term:        term,
		VoteGranted: voteGranted,
	}
}
