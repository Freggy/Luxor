package raft

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
}

// Start starts the raft node
func (n *Node) Start() error {
	if err := SetTimeoutConfig(n.MaxTimeout, n.MinTimeout); err != nil {
		return err
	}
	Timeout()
	// TODO: update state to candidate and ask for votes
	return nil
}
