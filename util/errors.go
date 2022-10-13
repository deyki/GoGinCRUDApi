package util

import "net/http"

type ErrorMessage struct {
	HttpStatus int
	Message    string
}

func (e ErrorMessage) FailedToOpenDB() *ErrorMessage {

	return &ErrorMessage{
		HttpStatus: 	http.StatusInternalServerError,
		Message: 		"Failed to open database!",
	}
}

func (e ErrorMessage) UserNotFound() *ErrorMessage {

	return &ErrorMessage{
		HttpStatus: 	http.StatusNotFound,
		Message: 		"User not found!",
	}
}