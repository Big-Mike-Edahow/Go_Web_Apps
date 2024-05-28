/* main.go */

package main

import (
	"html/template"
	"log"
	"net/http"
)

type Book struct {
	Id     int
	Year   string
	Author string
	Title  string
}

type BookList []Book

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
	myBookList := BookList{
		{1, "1982", "David Eddings", "Pawn of Prophecy"},
		{2, "1982", "David Eddings", "Queen of Sorcery"},
		{3, "1983", "David Eddings", "Magician's Gambit"},
		{4, "1984", "David Eddings", "Castle of Wizardry"},
		{5, "1984", "David Eddings", "Enchanter's End Game"},
	}

	indexTemplate.Execute(w, myBookList)
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
