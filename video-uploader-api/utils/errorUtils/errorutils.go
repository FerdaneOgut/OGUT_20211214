package errorutils

import "net/http"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewNotFoundError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
func NewUnprocessableEntityError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
func NewStatusConflictError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusConflict,
	}
}
func NewBadRequestError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}
