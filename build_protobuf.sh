#!/bin/bash

cd $(pwd)/pkg/raft/protocol
protoc --go_out=gen/ *.proto
