/* handlers.go */

package main

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("./templates/index.html")
	snippet := Snippet{}
	snippetArrays := []Snippet{}

	row, err := conn.Query("SELECT * FROM snippets")
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var id int
		var title, firstline, secondline, thirdline, author, created string
		err = row.Scan(&id, &title, &firstline, &secondline, &thirdline, &author, &created)
		if err != nil {
			panic(err.Error())
		}

		snippet.Id = id
		snippet.Title = title
		snippet.FirstLine = firstline
		snippet.SecondLine = secondline
		snippet.ThirdLine = thirdline
		snippet.Author = author
		snippet.Created = created

		snippetArrays = append(snippetArrays, snippet)
	}
	indexTemplate.Execute(w, snippetArrays)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	viewTemplate, _ := template.ParseFiles("./templates/view.html")
	id := r.URL.Query().Get("id")
	snippet := Snippet{}

	row, err := conn.Query("SELECT id, title, firstline, secondline, thirdline, author, created FROM snippets where id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var id int
		var title, firstline, secondline, thirdline, author, created string
		err = row.Scan(&id, &title, &firstline, &secondline, &thirdline, &author, &created)
		if err != nil {
			panic(err.Error())
		}

		snippet.Id = id
		snippet.Title = title
		snippet.FirstLine = firstline
		snippet.SecondLine = secondline
		snippet.ThirdLine = thirdline
		snippet.Author = author
		snippet.Created = created
	}
	viewTemplate.Execute(w, snippet)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	createTemplate := template.Must(template.ParseFiles("./templates/create.html"))
	createTemplate.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		firstline := r.FormValue("firstline")
		secondline := r.FormValue("secondline")
		thirdline := r.FormValue("thirdline")
		author := r.FormValue("author")

		if title == "" || firstline == "" || author == "" {
			http.Redirect(w, r, "/create", http.StatusMovedPermanently)
		}

		_, err := conn.Exec("INSERT INTO snippets (title, firstline, secondline, thirdline, author) VALUES (?, ?, ?, ?, ?)", title, firstline, secondline, thirdline, author)
		if err != nil {
			panic(err.Error())
		}
	
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	editTemplate := template.Must(template.ParseFiles("./templates/edit.html"))
	id := r.URL.Query().Get("id")
	snippet := Snippet{}

	row, err := conn.Query("SELECT id, title, firstline, secondline, thirdline, author FROM snippets WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var id int
		var title, firstline, secondline, thirdline, author string
		err = row.Scan(&id, &title, &firstline, &secondline, &thirdline, &author)
		if err != nil {
			panic(err.Error())
		}

		snippet.Id = id
		snippet.Title = title
		snippet.FirstLine = firstline
		snippet.SecondLine = secondline
		snippet.ThirdLine = thirdline
		snippet.Author = author
	}
	editTemplate.Execute(w, snippet)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		title := r.FormValue("title")
		firstline := r.FormValue("firstline")
		secondline := r.FormValue("secondline")
		thirdline := r.FormValue("thirdline")
		author := r.FormValue("author")

		if title == "" || firstline == "" || author == "" {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		_, err := conn.Exec("UPDATE snippets set title=?, firstline=?, secondline=?, thirdline=?, author=? where id=?", title, firstline, secondline, thirdline, author, id)
		if err != nil {
			panic(err.Error())
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := conn.Exec("DELETE FROM snippets WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}
