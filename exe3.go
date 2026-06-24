package main

import (
	"fmt"
	"net/http"
)

// r.Header.Get(), return an empty string.
func headersHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Custom-Token")
	if token == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(w, "Token received: %s", token)
	}
	cont := r.Header.Get("Content-Type")
	if cont == "" {
		r.Header[cont] = append(r.Header[cont], "Content-Type not provided")
	} else {
		fmt.Fprintf(w, "Content-Type: %s", cont)
	}
}
