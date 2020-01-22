package api

import (
	"net/http"
)

type ErrorCode int

const (
	unknown ErrorCode = iota
	BookNotFound
	BookCreationFailed
	BookUpdationFailed
	BookDeletionFailed
)

var statusCode = map[ErrorCode]int{
	BookNotFound:       http.StatusNotFound,
	BookCreationFailed: http.StatusNotModified,
	BookUpdationFailed: http.StatusNotModified,
	BookDeletionFailed: http.StatusNotFound,
}

func (e ErrorCode) HTTPStatus() int {
	if code, ok := statusCode[e]; ok {
		return code
	}
	return http.StatusInternalServerError
}
