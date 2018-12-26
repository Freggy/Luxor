package raft

import (
	"errors"
	"math/rand"
	"time"
)

var (
	minTimeOut int
	maxTimeOut int
)

// SetTimeoutConfig sets the minimum timeout and the maximum timeout.
func SetTimeoutConfig(max, min int) error {
	if max <= 0 || min <= 0 {
		return errors.New("values must be at least 1")
	}
	minTimeOut = min
	maxTimeOut = max
	return nil
}

// Timeout calls time#Sleep(duration) with a random generated duration
// in the range [minTimeout, maxTimeout)
func Timeout() {
	rand.Seed(time.Now().Unix())
	t := rand.Intn(maxTimeOut - minTimeOut) + minTimeOut
	time.Sleep(time.Duration(t) * time.Millisecond)
}
