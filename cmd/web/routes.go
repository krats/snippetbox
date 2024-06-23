package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Handle static assets
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// Display the home page
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	// Display a specific snippet
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	// Display a form for creating a new snippet
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	// Save a new snippet
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
