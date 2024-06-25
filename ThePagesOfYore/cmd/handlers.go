/* handlers.go */

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("./templates/index.html")

	books, err := getAllBooks()
	if err != nil {
		log.Println(err)
	}

	indexTemplate.Execute(w, books)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	viewTemplate := template.Must(template.ParseFiles("./templates/view.html"))

	bookId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(bookId)

	book, err := getOneBook(id)
	if err != nil {
		log.Println(err)
	}

	viewTemplate.Execute(w, book)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	addTemplate := template.Must(template.ParseFiles("./templates/add.html"))
	addTemplate.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		excerpt := r.FormValue("excerpt")
		bookPrice := r.FormValue("price")
		price, _ := strconv.ParseFloat(bookPrice, 32)

		if isbn == "" || title == "" || author == "" || excerpt == "" || price == 0 {
			http.Redirect(w, r, "/add", http.StatusMovedPermanently)
		}

		_, err := db.Exec("INSERT INTO books (isbn, title, author, excerpt, price) VALUES (?, ?, ?, ?, ?)", isbn, title, author, excerpt, price)
		if err != nil {
			log.Println(err)
		}

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	editTemplate := template.Must(template.ParseFiles("./templates/edit.html"))

	bookId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(bookId)

	book, err := getOneBook(id)
	if err != nil {
		log.Println(err)
	}

	editTemplate.Execute(w, book)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		excerpt := r.FormValue("excerpt")
		bookPrice := r.FormValue("price")
		price, _ := strconv.ParseFloat(bookPrice, 32)

		if isbn == "" || title == "" || author == "" || excerpt == "" || price == 0 {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		_, err := db.Exec("UPDATE books SET isbn=?, title=?, author=?, excerpt=?, price=? where id=?", isbn, title, author, excerpt, price, id)
		if err != nil {
			log.Println(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}
