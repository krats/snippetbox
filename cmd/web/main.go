package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Display the home page
	mux.HandleFunc("GET /{$}", home)
	// Display a specific snippet
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	// Display a form for creating a new snippet
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Save a new snippet
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	// Handle static assets
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	log.Print("Listening on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
