package config

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(message string, writer http.ResponseWriter) {
	type successdata struct {
		Message string `json:"message"`
	}
	temp := &successdata{Message: message}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

func ErrorResponse(error string, writer http.ResponseWriter) {
	type errdata struct {
		Message string `json:"message"`
	}
	temp := &errdata{Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}
