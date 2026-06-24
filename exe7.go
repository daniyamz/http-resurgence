package main

import (
	"html/template"
	"net/http"
)

const tmplstr = `
	<!DOCTYPE html>
	<html>
	<head><title>{{.Title}}</title></head>
	<body>
	<h1>{{.Title}}</h1>
	<p>{{.Body}}</p>
	{{if eq .Style "bold"}}<strong>{{.Body}}</strong>{{else}}{{.Body}}{{end}}
	</body>
	</html>
	`

type PageData struct {
	Title string
	Body  string
	Style string
}

func RenderHandler(w http.ResponseWriter, r *http.Request) {

	page := ""

	tmpl := template.Must(template.New(page).Parse(tmplstr))
	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")
	if title == "" || body == "" {
		http.Error(w, "title and body are required.", 400)
		return
	}
	//w.Header().Set("Content-Text", "text/plain")
	err := tmpl.Execute(w, PageData{Title: title, Body: body})
	if err != nil {
		http.Error(w, "template execution failed", 500)
		return
	}
}
