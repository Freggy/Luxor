package raft

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

var (
	minTimeOut int
	maxTimeOut int
	salt int64
)

// SetTimeoutConfig sets the minimum timeout and the maximum timeout.
func SetTimeoutConfig(max, min int, id uuid.UUID) error {
	if max <= 0 || min <= 0 {
		return errors.New("values must be at least 1")
	}
	sha := sha256.New()
	sha.Write(id.Bytes())

	salt = int64(binary.BigEndian.Uint64(sha.Sum(nil)))
	minTimeOut = min
	maxTimeOut = max
	return nil
}

// Timeout calls time#Sleep(duration) with a random generated duration
// in the range [minTimeout, maxTimeout)
func Timeout() time.Duration {
	rand.Seed(time.Now().Unix() + salt)
	t := rand.Intn(maxTimeOut - minTimeOut) + minTimeOut
	dur := time.Duration(t) * time.Millisecond
	time.Sleep(dur)
	return dur
}
