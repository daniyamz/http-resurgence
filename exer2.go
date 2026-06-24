package main

import (
	"fmt"
	"io"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, err := io.ReadAll(r.Body)
	for err != nil {
		fmt.Fprintln(w, "Error occured", err)
	}
	defer r.Body.Close()
	data := string(file)
	if data == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s", data)
}
