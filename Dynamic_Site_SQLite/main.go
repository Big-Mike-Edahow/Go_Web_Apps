/* main.go */

package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

var tmpl *template.Template

var db *sql.DB

func init() {
	tmpl = template.Must(template.ParseFiles("./templates/index.html", "./templates/about.html"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Serving HTTP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}

func getSQLiteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func queryDB(w http.ResponseWriter) {
	type Book struct {
		Id     int
		Year   string
		Author string
		Title  string
	}
	type ListOfBooks []Book
	myBookList := ListOfBooks{}
	novel := Book{}

	rows, _ := db.Query("SELECT * FROM books;")
	for rows.Next() {
		rows.Scan(&novel.Id, &novel.Year, &novel.Author, &novel.Title)
		myBookList = append(myBookList, novel)
	}
	tmpl.Execute(w, myBookList)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type Booklist struct {
		Id     string
		Year   string
		Author string
		Title  string
	}

	book := Booklist{
		Id:     r.FormValue("id"),
		Year:   r.FormValue("year"),
		Author: r.FormValue("author"),
		Title:  r.FormValue("title"),
	}

	db = getSQLiteDB()

	if r.Method != http.MethodPost {
		queryDB(w)
	}

	if r.FormValue("submit") == "Read" {
		queryDB(w)
	}

	if r.FormValue("submit") == "Insert" {
		Id, _ := strconv.Atoi(book.Id)
		_, err := db.Exec("INSERT INTO books (id, year, author, title) VALUES (?,?,?,?)", Id, book.Year, book.Author, book.Title)
		if err != nil {
			log.Println(err)
			queryDB(w)
		} else {
			queryDB(w)
		}
	}

	if r.FormValue("submit") == "Update" {
		Id, _ := strconv.Atoi(book.Id)
		_, err := db.Exec("UPDATE books SET year=?, author=?, title=? WHERE id=?", book.Year, book.Author, book.Title, Id)
		if err != nil {
			log.Println(err)
			queryDB(w)
		} else {
			queryDB(w)
		}

	}

	if r.FormValue("submit") == "Delete" {
		Id, _ := strconv.Atoi(book.Id)
		_, err := db.Exec("DELETE FROM books WHERE id=?", Id)
		if err != nil {
			log.Println(err)
			queryDB(w)
		} else {
			queryDB(w)
		}
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
