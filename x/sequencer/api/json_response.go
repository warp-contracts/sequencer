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

type httpJsonError struct {
	ErrorType    string     `json:"type"`
	ErrorMessage string     `json:"message"`
	Status       httpStatus `json:"status"`
}

// Writes a bad request error in the form of JSON to the HTTP response
func BadRequestError(w http.ResponseWriter, err error, errorType string) {
	jsonError := createJsonError(err.Error(), errorType, http.StatusBadRequest)
	writeError(w, jsonError)
}

// Writes a internal server error in the form of JSON to the HTTP response (takes an error as a string)
func InternalServerErrorString(w http.ResponseWriter, err string, errorType string) {
	jsonError := createJsonError(err, errorType, http.StatusInternalServerError)
	writeError(w, jsonError)
}

// Writes a internal server error in the form of JSON to the HTTP response
func InternalServerError(w http.ResponseWriter, err error, errorType string) {
	InternalServerErrorString(w, err.Error(), errorType)
}

// Returns a JSON response with a 404 status
func NotFoundResponse(w http.ResponseWriter, message string) {
	jsonError := createJsonError(message, "not found", http.StatusNotFound)
	writeError(w, jsonError)
}

func createJsonError(err string, errorType string, code int) httpJsonError {
	return httpJsonError{
		ErrorType:    errorType,
		ErrorMessage: err,
		Status: httpStatus{
			Code: code,
			Text: http.StatusText(code),
		},
	}
}

func writeError(w http.ResponseWriter, jsonError httpJsonError) {
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