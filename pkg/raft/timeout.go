package raft

import (
	"crypto"
	"encoding/binary"
	"errors"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

var (
	minTimeOut int
	maxTimeOut int
	idChecksum uint64
)

// SetTimeoutConfig sets the minimum timeout and the maximum timeout.
func SetTimeoutConfig(max, min int, id uuid.UUID) error {
	if max <= 0 || min <= 0 {
		return errors.New("values must be at least 1")
	}

	sha := crypto.SHA256.New()
	sha.Write(id.Bytes())

	idChecksum = binary.BigEndian.Uint64(sha.Sum(nil))
	minTimeOut = min
	maxTimeOut = max
	return nil
}

// Timeout calls time#Sleep(duration) with a random generated duration
// in the range [minTimeout, maxTimeout)
func Timeout() {
	rand.Seed(time.Now().Unix() + int64(idChecksum))
	t := rand.Intn(maxTimeOut - minTimeOut) + minTimeOut
	time.Sleep(time.Duration(t) * time.Millisecond)
}
