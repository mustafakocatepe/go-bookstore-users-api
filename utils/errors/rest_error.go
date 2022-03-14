package errors

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		message,
		http.StatusBadRequest,
		"bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		message,
		http.StatusNotFound,
		"not_found",
	}
}
