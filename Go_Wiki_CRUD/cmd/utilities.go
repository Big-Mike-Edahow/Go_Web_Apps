/* utilities.go */

package main

import (
	"log"
	"net/http"
	"os"
)

func (p *Page) save() error {
	dataDir := "./data/"
	filename := p.Title
	return os.WriteFile(dataDir+filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	dataDir := "./data/"
	filename := title
	body, err := os.ReadFile(dataDir + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
