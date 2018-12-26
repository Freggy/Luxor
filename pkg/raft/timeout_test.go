package raft

import (
	"fmt"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/out"
	"testing"
)

func Test_SetTimeoutConfig_Successful(t *testing.T) {
	if err := SetTimeoutConfig(2, 3); err != nil {
		t.Fail()
	}
}

func Test_SetTimeoutConfig_With_Smaller_Values(t *testing.T) {
	/*
	if err := SetTimeoutConfig(0, -3); err == nil {
		t.Fail()
	}*/

	var l interface{}
	l = raft.AppendEntriesRequest{}

	switch v := l.(raft.AppendEntriesRequest) {
	default:
		fmt.Printf("unexpected type %T", v)
	case raft.AppendEntriesRequest:
		fmt.Println("lol")
	}
	
}
