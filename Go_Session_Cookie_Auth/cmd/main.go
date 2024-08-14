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
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	initDB()

	log.Println("Serving HTTP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}
