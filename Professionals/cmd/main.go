/* main.go */

package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Prof struct {
	Id int
	Name string
	Age string
	Employ string
	Created string
}

var conn, _ = sql.Open("sqlite3", "./data/database.db")

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Listening and serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}
