package main

import (
	"fmt"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if "Content-Type" != r.Header.Get("application/x-www-form-urlencoded") {
		http.Error(w, "Unsupported Media", http.StatusUnsupportedMediaType)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "can't parse", 404)
		return
	}
	uname := r.FormValue("username")
	lang := r.FormValue("language")

	if uname == "" || lang == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		if uname == "" {
			fmt.Fprintln(w, "username is required")
		}
		if lang == "" {
			fmt.Fprintln(w, "language is required")
		}
	} else {
		fmt.Fprintf(w, "Hello %s, you are coding in %s!", uname, lang)

	}
}
