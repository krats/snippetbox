package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	// Display the home page
	mux.HandleFunc("GET /{$}", app.home)
	// Display a specific snippet
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	// Display a form for creating a new snippet
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	// Save a new snippet
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	// Handle static assets
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	return mux
}
