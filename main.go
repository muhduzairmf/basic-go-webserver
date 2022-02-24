package main

import (
	"log"
	"net/http"
	"basic-go-webserver/handler"
	// Handler package for managing each routes
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// w variable is for response to the client
	// r variable is for receive the request from client

	if r.URL.Path != "/" {
		w.Write([]byte("<h1 style='color: red;'>NOT FOUND</h1>"))
		// or
		// http.NotFound(w, r)
		return
	}
	// Please note that anything inside the home path "/"
	// will also executed if the path is not match
	// to the defined path at main function.
	// So that, this if statement will filter the invalid path.

	// w.Write([]byte("HELLO WORLD!!!"))
	// If we use plain text like this it will response
	// Content-Type as a text/plain

	w.Write([]byte("<h1>HELLO WORLD!!!</h1>"))
	// If we use some html tag, it will auto response
	// Content-Type as a text/html
	
	// Other that w.Write() function, we also can use
	// fmt.Fprintln(w, "Hello World")
	// or
	// fmt.Fprintln(w, "<h1>Hello World</h1>")
	// or other variation like Fprint, Fprintf etc...
}

func main() {
	var port = ":8080"
	// Set up the port

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>About page</h1>"))
	}
	// Variable as a function

	http.HandleFunc("/", HomeHandler)
	// The path/route and function handler

	http.HandleFunc("/products", handler.ProductHandler)
	// Import the product handler from handler package

	http.HandleFunc("/features", handler.FeaturesHandler)

	http.HandleFunc("/article", handler.ArticleHandler)

	http.HandleFunc("/feedback", handler.FeedbackHandler)
	
	http.HandleFunc("/about", aboutHandler)

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1'>Profile</h1>"))
	})
	// Function as a parameter

	http.HandleFunc("/api/feedback", handler.APIFeedbackHandler)
	// This is the route/path for API instead of webpage

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	// This is how to host a static assets like css, js or images
	// http.Dir() function is include all the static item 
	// in the static folder to be hosted.
	// For example, the path to get css file is at
	// http://localhost:8080/static/style.css

	log.Println("Server is listening on http://localhost" + port)
	http.ListenAndServe(port, nil)
	// Open and set up the server port
}
