/* main.go */

package main

import (
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/savefile", saveFileHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Listening and serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}
