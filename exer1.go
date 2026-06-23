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
	http.ListenAndServe(":8080", mux)
}
