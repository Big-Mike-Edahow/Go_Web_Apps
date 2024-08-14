/* handlers.go */

package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string
	Username string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("./templates/index.html")
	indexTemplate.Execute(w, nil)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}

	/* Salt and hash the password using the bcrypt algorithm. The
	second argument is the computing power you wish to utilize. */
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		log.Println(err)
	}

	// Insert the username and hashed password into the database.
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", creds.Username, string(hashedPassword))
	if err != nil {
		log.Println(err)
	}

	welcomeTemplate, _ := template.ParseFiles("./templates/welcome.html")
	welcomeTemplate.Execute(w, nil)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}

	// Get the existing entry present in the database for the given username
	result := db.QueryRow("SELECT password FROM users WHERE username=?", creds.Username)
	// Create an instance of `Credentials` to store values from the database.
	storedCreds := &Credentials{}
	// Store the obtained password in `storedCreds`
	err := result.Scan(&storedCreds.Password)
	if err != nil {
		log.Println(err)
	}

	// Compare the stored hashed password, with the hashed password that was received.
	err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password))
	if err != nil {
		// If username or password doesn't match display error page.
		errorTemplate, _ := template.ParseFiles("./templates/error.html")
		errorTemplate.Execute(w, nil)
	} else {
		// If username and password match, display welcome page.
		welcomeTemplate, _ := template.ParseFiles("./templates/welcome.html")
		welcomeTemplate.Execute(w, nil)
	}
}
