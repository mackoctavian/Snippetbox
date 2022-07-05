package main

import (
	"log"
	"net/http"
)

func main() {
	//Initilizing a http.NewServeMux() func and registering handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Using http.ListenAndServe to start a new server passing two parameters: The TCP address to listen and the created serveMux
	err := http.ListenAndServe(":4000", mux)

	//Checking if the web server returns an error and using log.fatal to log the error and exit
	if err != nil {
		log.Fatal(err)
	}
}
