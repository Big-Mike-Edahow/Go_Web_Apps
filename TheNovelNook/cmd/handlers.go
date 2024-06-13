/* handlers.go */

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// booksIndex sends a HTTP response listing all books.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("./templates/index.html")
	books, err := AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	indexTemplate.Execute(w, books)
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
		bookPrice := r.FormValue("price")
		price, _ := strconv.ParseFloat(bookPrice, 32)

		if isbn == "" || title == "" || author == "" || price == 0 {
			http.Redirect(w, r, "/add", http.StatusMovedPermanently)
		}

		_, err := DB.Exec("INSERT INTO books (isbn, title, author, price) VALUES (?, ?, ?, ?)", isbn, title, author, price)
		if err != nil {
			panic(err.Error())
		}

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	editTemplate := template.Must(template.ParseFiles("./templates/edit.html"))
	id := r.URL.Query().Get("id")
	book := Book{}

	row, err := DB.Query("SELECT id, isbn, title, author, price FROM books WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var id int
		var isbn, title, author string
		var price float32
		err = row.Scan(&id, &isbn, &title, &author, &price)
		if err != nil {
			panic(err.Error())
		}

		book.Id = id
		book.Isbn = isbn
		book.Title = title
		book.Author = author
		book.Price = price
	}
	editTemplate.Execute(w, book)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		bookPrice := r.FormValue("price")
		price, _ := strconv.ParseFloat(bookPrice, 32)

		if isbn == "" || title == "" || author == "" || price == 0 {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		_, err := DB.Exec("UPDATE books SET isbn=?, title=?, author=?, price=? where id=?", isbn, title, author, price, id)
		if err != nil {
			panic(err.Error())
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}
