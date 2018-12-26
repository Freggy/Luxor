package logging

import (
	"github.com/luxordynamics/luxor/pkg/raft"
	"testing"
)

func Test_SetTimeoutConfig_Successful(t *testing.T) {
	if err := raft.SetTimeoutConfig(2, 3); err != nil {
		t.Fail()
	}
}

func Test_SetTimeoutConfig_With_Smaller_Values(t *testing.T) {
	if err := raft.SetTimeoutConfig(0, -3); err == nil {
		t.Fail()
	}
}
