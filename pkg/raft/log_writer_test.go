package raft

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/gen"
	"log"
	"testing"
)

func TestBinaryTextWriter_Write(t *testing.T) {
	w := BinaryLogEntryWriter{}

	payload1 := createPlayload(4, true)
	payload2 := createPlayload(55, false)
	payload3 := createPlayload(1337, false)

	entries := []*raft.Entry{
		payload1,
		payload2,
		payload3,
	}

	ba, err := w.Write(entries)

	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(ba)

	resp1, err := readPlayload(buf)

	if err != nil {
		log.Fatal(err)
	}

	if resp1.GetVoteGranted() != true && resp1.GetTerm() != 4 {
		t.Fail()
	}

	// skip one playload
	readPlayload(buf)

	resp3, err := readPlayload(buf)

	if err != nil {
		log.Fatal(err)
	}

	if resp3.GetVoteGranted() != true && resp3.GetTerm() != 1337 {
		t.Fail()
	}

}

func createPlayload(term uint32, bool2 bool) *raft.Entry {
	payload := &raft.VoteResponse{Term: term, VoteGranted: bool2}
	bArr, _ := proto.Marshal(payload)
	return &raft.Entry{Payload: bArr}
}

func readPlayload(buf *bytes.Buffer) (*raft.VoteResponse, error) {
	var length uint32
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return nil, err
	}

	b := make([]byte, length)

	if _, err := buf.Read(b); err != nil {
		return nil, err
	}

	var data raft.VoteResponse

	if err := proto.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
