package raft

import "github.com/satori/go.uuid"

const (
	Follower State = iota
	Candidate
	Leader
)

type State int

type Node struct {
	MinTimeout int
	MaxTimeout int
	State State
	Uuid uuid.UUID
}

// Start starts the raft node
func (n *Node) Start() error {
	n.Uuid = uuid.NewV4()
	if err := SetTimeoutConfig(n.MaxTimeout, n.MinTimeout, n.Uuid); err != nil {
		return err
	}
	Timeout()
	// TODO: update state to candidate and ask for votes
	return nil
}

func (n *Node) ChangeState(state State) {
	if state == Follower {
		n.follow()
	} else if state == Candidate {
		n.candidate()
	} else if state == Leader {
		n.lead()
	}
}

func (n *Node) follow() {

}

func (n *Node) candidate() {

}

func (n *Node) lead() {

}
