package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//Defining a home handler func which writes  byte of slice containg "Hello from snippet box" as a response body
func home(w http.ResponseWriter, r *http.Request) {
	//checking if the url pattern matches "/" and sending 404 code if it doesn't
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//Response body
	w.Write([]byte("Hello from snippetbox"))
}

//Adding a showSnippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {
	//Extracting the value of the id parameter from the query string and converting it to integer
	//using strconv.Atoi() func
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	//if the value is less than 1, We return a 404 page not found response
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//Using the fmt.Fprintf() func to interpolate the id value with our response
	fmt.Fprintf(w, "Display a specifi snippet with ID %d...", id)
}

//Adding a createSnippet handler
func createSnippet(w http.ResponseWriter, r *http.Request) {
	//Using r.method to check whether the request is using a POST
	//and sending 405 status code if it's not
	if r.Method != "POST" {
		//Using the Header().Set() method to add an 'Allow: POST' header to the response header map
		//The first parameter is the header name and the second is the header value
		w.Header().Set("Allow", "POST")

		//Using http.Error() func 405 status code and method not allowed as response body
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
