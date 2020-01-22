package api

import "fmt"

type BookError struct {
	Code        ErrorCode `json:"code,omitempty"`
	Message     string    `json:"message,omitempty"`
	Description string    `json:"description,omitempty"`
}

func (b BookError) Error() string {
	return fmt.Sprintf("code: %d msg: %s ", b.Code, b.Description)
}
