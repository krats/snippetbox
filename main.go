package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello from Snippetbox!"))
	log.Fatal(err)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Print("Listening on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
