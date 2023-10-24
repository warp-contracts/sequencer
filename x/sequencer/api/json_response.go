package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpStatus struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type httpJsonError[T any] struct {
	ErrorType    string     `json:"type"`
	ErrorMessage T          `json:"message"`
	Status       httpStatus `json:"status"`
}

// Writes a bad request error in the form of JSON to the HTTP response
func BadRequestError(w http.ResponseWriter, err error, errorType string) {
	ErrorWithStatus(w, err.Error(), errorType, http.StatusBadRequest)
}

// Writes a internal server error in the form of JSON to the HTTP response (takes an error as a string)
func InternalServerErrorString(w http.ResponseWriter, message string, errorType string) {
	ErrorWithStatus(w, message, errorType, http.StatusInternalServerError)
}

// Writes a internal server error in the form of JSON to the HTTP response
func InternalServerError(w http.ResponseWriter, err error, errorType string) {
	InternalServerErrorString(w, err.Error(), errorType)
}

// Returns a JSON response with a 404 status
func NotFoundResponse(w http.ResponseWriter, message string) {
	ErrorWithStatus(w, message, "not found", http.StatusNotFound)
}

func ErrorWithStatus[T any](w http.ResponseWriter, message T, errorType string, statusCode int) {
	jsonError := createJsonError(message, errorType, statusCode)
	writeError(w, jsonError)
}

func createJsonError[T any](errorMessage T, errorType string, statusCode int) httpJsonError[T] {
	return httpJsonError[T]{
		ErrorType:    errorType,
		ErrorMessage: errorMessage,
		Status: httpStatus{
			Code: statusCode,
			Text: http.StatusText(statusCode),
		},
	}
}

func writeError[T any](w http.ResponseWriter, jsonError httpJsonError[T]) {
	setHeaders(w)
	w.WriteHeader(jsonError.Status.Code)
	if err := json.NewEncoder(w).Encode(jsonError); err != nil {
		panic(err)
	}
}

// Returns a response with a 200 status
// Encodes the provided content into JSON format
func OkResponse(w http.ResponseWriter, response any) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		InternalServerError(w, err, "response encoding error")
		return
	}

	_, err = fmt.Fprintf(w, "%s", jsonResponse)
	if err != nil {
		InternalServerError(w, err, "response writing error")
		return
	}
	setHeaders(w)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}
