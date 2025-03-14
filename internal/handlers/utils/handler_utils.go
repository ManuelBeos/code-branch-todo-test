package handler_utils

import (
	"encoding/json"
	"net/http"

	"github.com/manuelbeos/code-branch-todo-test/internal/handlers/dtos"
)

func HandlerErrorResponse(rw http.ResponseWriter, statusCode int, err error) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err.(type) {
	case *dtos.ErrorResponse:
		errorResponse := err.(*dtos.ErrorResponse)
		jsonData, err := json.Marshal(errorResponse)
		if err != nil {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(statusCode)
		_, err = rw.Write(jsonData)
		if err != nil {
			return
		}

	default:
		rw.WriteHeader(statusCode)
		_, err = rw.Write([]byte(err.Error()))
		if err != nil {
			return
		}

	}
}

func HandlerSuccessResponse(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(statusCode)
	_, err = rw.Write(jsonData)
	if err != nil {
		return
	}
}
