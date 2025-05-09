package rest_error

import "net/http"

type RestError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: message,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Code:   http.StatusNotFound,
		Status: "Not Found",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: message,
	}
}
