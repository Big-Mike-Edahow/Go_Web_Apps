// main.go

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/view/{id}", viewHandler)
	mux.HandleFunc("/create", createHandler)
	mux.HandleFunc("/about", aboutHandler)

	log.Printf("Starting HTTP Server on port %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, logRequest(mux)))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./templates/base.html",
		"./templates/index.html",
	}

	indexTemplate, _ := template.ParseFiles(files...)
	indexTemplate.ExecuteTemplate(w, "base", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./templates/base.html",
		"./templates/view.html",
	}

	viewTemplate, _ := template.ParseFiles(files...)
	viewTemplate.ExecuteTemplate(w, "base", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./templates/base.html",
		"./templates/create.html",
	}

	createTemplate, _ := template.ParseFiles(files...)
	createTemplate.ExecuteTemplate(w, "base", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/base.html",
		"./templates/about.html",
	}

	aboutTemplate, _ := template.ParseFiles(files...)
	aboutTemplate.ExecuteTemplate(w, "base", nil)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
