package server

import (
	"encoding/json"
	"net/http"
)

// This file contains functions to be used frequently for HTTP response
// delivery.

func sendResponse(w http.ResponseWriter, msg string, code int) error {

	encoded, err := json.Marshal(&genericResponse{Message: msg})
	if err != nil {
		return err
	}
	w.WriteHeader(code)
	_, err = w.Write(encoded)
	if err != nil {
		return err
	}
	return nil
}

func sendServerError(w http.ResponseWriter) error {
	return sendResponse(w, "server error", http.StatusInternalServerError)
}
