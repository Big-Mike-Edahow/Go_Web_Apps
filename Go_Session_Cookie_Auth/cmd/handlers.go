/* handlers.go */

package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// Username and Password login credentials
type Credentials struct {
	Password string
	Username string
}

// Each session contains username and expiry time.
type Session struct {
	Username string
	Expiry   time.Time
}

// This map stores the users sessions.
var sessions = map[string]Session{}

// we'll use this method later to determine if the session has expired
func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("./templates/index.html")
	indexTemplate.Execute(w, nil)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}
	result := db.QueryRow("SELECT password FROM users WHERE username=?", creds.Username)
	// Create an instance of `Credentials` to store values from the database.
	storedCreds := &Credentials{}
	// Store the obtained password in `storedCreds`
	err := result.Scan(&storedCreds.Password)
	if err != nil {
		log.Println(err)
	}
	if storedCreds.Password == creds.Password {
		// Create a new random session token
		sessionToken := uuid.NewString()
		expiresAt := time.Now().Add(120 * time.Second)
		// Set the token in the session map, along with the user whom it represents.
		sessions[sessionToken] = Session{
			Username: creds.Username,
			Expiry:   expiresAt,
		}

		/* Finally, we set the client cookie for "session_token" as the session token
		   we just generated we also set an expiry time of 120 seconds. */
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiresAt,
		})
		// Redirect to a welcome or login page after signup.
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		// Redirect to an error page.
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// We then get the name of the user from our session map, where we set the session token
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	welcomeTemplate, _ := template.ParseFiles("./templates/welcome.html")
	welcomeTemplate.Execute(w, userSession.Username)
	// Finally, return the welcome message to the user
	// w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.Username)))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	errorTemplate, _ := template.ParseFiles("./templates/error.html")
	errorTemplate.Execute(w, nil)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code from this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// If the previous session is valid, create a new session token for the current user
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	sessions[newSessionToken] = Session{
		Username: userSession.Username,
		Expiry:   expiresAt,
	}

	// Delete the older session token
	delete(sessions, sessionToken)

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// remove the users session from the session map
	delete(sessions, sessionToken)

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
	LoggedOutTemplate, _ := template.ParseFiles("./templates/logged_out.html")
	LoggedOutTemplate.Execute(w, nil)
}
