// main.go

package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id        int
	Completed int
	Item      string
}

var conn, _ = sql.Open("sqlite3", "./data/database.db")

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/complete", completeHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Listening and serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))

	defer conn.Close()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("./templates/index.html"))
	todo := Todo{}
	todoArrays := []Todo{}

	rows, err := conn.Query("SELECT * FROM todos")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var completed int
		var item string
		err = rows.Scan(&id, &completed, &item)
		if err != nil {
			panic(err.Error())
		}
		todo.Id = id
		todo.Item = item
		todo.Completed = completed

		todoArrays = append(todoArrays, todo)
	}
	indexTemplate.Execute(w, todoArrays)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		item := r.FormValue("item")
		completed := 0

		if item == "" {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		_, err := conn.Exec("INSERT INTO todos (item, completed) VALUES (?, ?)", item, completed)
		if err != nil {
			panic(err.Error())
		}

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func completeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := conn.Exec("UPDATE todos SET completed = 1 WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := conn.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
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
