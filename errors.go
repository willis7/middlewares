package middlewares

import (
	"net/http"
	"log"
	"encoding/json"
)

type(
	appError struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

// DisplayError is a helper function which formats and writes an error in JSON format
func DisplayError(w http.ResponseWriter, handleError error, message string, code int) {
	errObj := appError{
		Error: handleError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("[App Error]: %s\n", handleError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}
