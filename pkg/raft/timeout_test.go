package raft

import (
	"github.com/satori/go.uuid"
	"testing"
)

func Test_SetTimeoutConfig_Successful(t *testing.T) {
	if err := SetTimeoutConfig(300, 200, uuid.NewV4()); err != nil {
		t.Fail()
	}
}

func Test_SetTimeoutConfig_With_Smaller_Values(t *testing.T) {
	if err := SetTimeoutConfig(0, -3, uuid.NewV4()); err == nil {
		t.Fail()
	}
}
