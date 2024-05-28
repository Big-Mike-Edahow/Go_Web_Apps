// main.go
/* This example will show how to store data in session cookies using the
   popular gorilla/sessions package in Go. Cookies are small pieces of data
   stored in the browser of a user and are sent to our server on each request. */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

func main() {
	http.HandleFunc("/secret", secretHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	log.Println("Serving HTTP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
