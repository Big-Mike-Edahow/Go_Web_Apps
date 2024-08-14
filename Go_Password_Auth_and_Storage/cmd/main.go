/* main.go */

package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signin", signInHandler)
	http.HandleFunc("/signup", signUpHandler)

	initDB()

	log.Println("Serving HTTP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}
