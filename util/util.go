// Package util provides utility functions for goresource.
package util

import (
	"encoding/json"
	"net/http"
)

// WriteJSON marshals the given data into json and writes it to the response stream.
func WriteJSON(data interface{}, rw http.ResponseWriter) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonBytes)
}
