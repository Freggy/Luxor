package rest

import "encoding/json"

// Struct representing a error that will be sent to the client.
type Error struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	ErrorType string `json:"type"`
}

// NewError creates a new Error type
func NewError(code int, message string, errorType string) Error {
	return Error{code, message, errorType}
}

// ToJson marshals the struct to JSON by using
//  json.Marshal(e)
func (e Error) ToJson() ([]byte, error) {
	return json.Marshal(e)
}
