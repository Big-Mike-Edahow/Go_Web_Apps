/* helpers.go */

package main

import (
	"log"
)

func getOneBook(id int) (Book, error) {
	row, err := db.Query("SELECT id, isbn, title, author, excerpt, price FROM books WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	var book Book
	for row.Next() {
		err = row.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Excerpt, &book.Price)
		if err != nil {
			log.Println(err)
		}
	}
	return book, err
}

func getAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		
		err := rows.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Excerpt, &book.Price)
		if err != nil {
			log.Println(err)
		}
		books = append(books, book)
	}
	return books, nil
}
