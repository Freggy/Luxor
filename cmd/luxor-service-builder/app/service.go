package app

import (
	"encoding/json"
	"io/ioutil"
)

type Service struct {
	Id    string    `json:"id"`
	Files []FileDef `json:"files"`
}

type FileDef struct {
	Path  string    `json:"path"`
	Files []FileDef `json:"files,omitempty"`
}

func NewServiceFromFile(path string) (*Service, error) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	service, err := NewServiceFromJson(bytes)

	if err != nil {
		return nil, err
	}

	return service, nil
}

func NewServiceFromJson(data []byte) (*Service, error) {
	var s Service

	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
