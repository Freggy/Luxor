package raft

import (
	"bytes"
	"encoding/binary"
	"github.com/luxordynamics/luxor/pkg/raft/protocol/gen"
)

type LogEntryWriter interface {
	Write([]*raft.Entry) ([]byte, error)
}

type BinaryLogEntryWriter struct {

}

func (b *BinaryLogEntryWriter) Write(log []*raft.Entry) ([]byte, error) {
	buf := new(bytes.Buffer)

	for i := 0; i < len(log); i++ {
		entry := log[i]

		if err := binary.Write(buf, binary.BigEndian, uint32(len(entry.GetPayload()))); err != nil {
			return nil, err
		}

		if _, err := buf.Write(entry.GetPayload()); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
