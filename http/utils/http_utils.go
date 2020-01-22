package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ASinha24/LibraryManagementSystem/api"
)

func WriteResponse(status int, response interface{}, rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	if response != nil {
		json.NewEncoder(rw).Encode(response)
	}

}

func WriteErrorResponse(status int, err error, rw http.ResponseWriter) {
	bookErr, ok := err.(*api.BookError)
	if !ok {
		bookErr = &api.BookError{
			Code:        0,
			Message:     "failed in serving request",
			Description: bookErr.Error(),
		}
	} else {
		status = bookErr.Code.HTTPStatus()
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(bookErr)

}
