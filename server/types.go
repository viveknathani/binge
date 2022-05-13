package server

// This file defines how the incoming and outgoing JSON payloads look like.

type genericResponse struct {
	Message string `json:"message"`
}
