package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func StatusCodeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code parament is required", 400)
		return
	}
	digit, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid integer", 400)
		return
	}
	if digit <= 100 && digit >= 599 {
		http.Error(w, "code must be a valid HTTP status code (100-599)", 400)
		return
	}
	w.WriteHeader(404)
	fmt.Fprintf(w, "Responding with status %s", code)

}
