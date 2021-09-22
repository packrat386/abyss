package main

import (
	"encoding/xml"
	"net/http"
)

type ErrorResponse struct {
	Error Error `xml:"Error"`
}

type Error struct {
	Code      string `xml:"Code"`
	Message   string `xml:"Message"`
	Resource  string `xml:"Resource"`
	RequestID string `xml:"RequestId"`
}

func renderError(w http.ResponseWriter, scode int, code, message, resource, requestID string) {
	e := ErrorResponse{
		Error: Error{
			Code:      code,
			Message:   message,
			Resource:  resource,
			RequestID: requestID,
		},
	}

	w.Header().Add("Content-Type", "application/xml")

	w.WriteHeader(scode)
	w.Write([]byte(xml.Header))

	xml.NewEncoder(w).Encode(e)
}
