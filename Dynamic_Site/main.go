/* main.go */

package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Listening and serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("./templates/index.html"))
	indexTemplate.Execute(w, nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
