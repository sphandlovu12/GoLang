package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Initialize your URL mapping data structure here.

	http.HandleFunc("/", handleShorten)
	http.HandleFunc("/short", handleRedirect)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Implement URL shortening logic and respond with the shortened URL.
	} else {
		// Serve a web page where users can enter a long URL.
		// You can use the html/template package to create the page.
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	// Implement URL redirection logic here.
}
