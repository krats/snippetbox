package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox!"))
	if err != nil {
		log.Print(err)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// id should be an integer greater than zero.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("Snippet view for id : %d", id)))
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
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Listening on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
