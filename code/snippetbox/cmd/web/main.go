package main

import (
	"log"
	"net/http"
)

func main() {
	//initialize new servemux
	//register home function as handler for "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Start new web server using Listenandserve
	//return error using log.fatal
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
