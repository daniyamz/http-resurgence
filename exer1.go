package main

import (
	"fmt"
	"net/http"
)

func inspectoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You made a %s request\n", r.Method)
}

func main() {
	fmt.Println("Server is runing on port :8080, Press ctrl + c to ternimate the server")
	mux := http.NewServeMux()
	mux.HandleFunc("/method-inspector", inspectoHandler)
	mux.HandleFunc("/echo", EchoHandler)
	mux.HandleFunc("/headers", headersHandler)
	mux.HandleFunc("/form", FormHandler)
	mux.HandleFunc("/status", StatusCodeHandler)
	http.ListenAndServe(":8080", mux)
}
