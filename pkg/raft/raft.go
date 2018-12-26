package raft

import (
	"math/rand"
	"time"
)

type Timeout struct {
	timeout int
}

func NewTimeout(min, max int) *Timeout {
	rand.Seed(time.Now().Unix())
	return &Timeout{
		timeout: rand.Intn(max - min) + min,
	}
}

func (t Timeout) DoTimeout()  {
	time.Sleep(time.Duration(t.timeout) * time.Millisecond)
}
