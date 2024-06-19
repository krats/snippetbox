package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox!"))
	if err != nil {
		log.Print(err)
	}
}

func snippetView(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Snippet view!"))
	if err != nil {
		log.Print(err)
	}
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Snippet create!"))
	if err != nil {
		log.Print(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Listening on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
