/* main.go */

package main

import (
	"database/sql"
    "log"
    "net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Snippet struct {
    Id      int
    Title   string
	FirstLine string
	SecondLine string
	ThirdLine string
	Author string
    Created string
}

var conn, _ = sql.Open("sqlite3", "./data/database.db")

func main() {
    fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/about", aboutHandler)
    
    log.Println("Listening and serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))

	defer conn.Close()
}
