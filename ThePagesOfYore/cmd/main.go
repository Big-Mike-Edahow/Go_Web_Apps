/* main.go */

package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Id      int
	Isbn    string
	Title   string
	Author  string
	Excerpt string
	Price   float32
}

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("db connection failed")
		panic(err)
	}
	log.Println("Connected to the database.")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Serving HTTP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}
