package dtos

import "strconv"

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *ErrorResponse) Error() string {
	return e.Message + " - " + strconv.Itoa(e.Code)
}

func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    code,
	}
}
