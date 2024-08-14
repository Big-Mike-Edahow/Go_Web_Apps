/* handlers.go */

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAllTodos()
	if err != nil {
		log.Println(err)
	}

	indexTemplate := template.Must(template.ParseFiles("./templates/index.html"))
	indexTemplate.Execute(w, todos)
}

func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		createTemplate := template.Must(template.ParseFiles("./templates/create.html"))
		createTemplate.Execute(w, nil)
	case "POST":
		item := r.FormValue("item")
		completed := 0

		msg := &Message{
			Item: r.PostFormValue("item"),
		}

		if !msg.Validate() {
			createTemplate := template.Must(template.ParseFiles("./templates/create.html"))
			createTemplate.Execute(w, msg)
		} else {
			err := app.todos.Insert(item, completed)
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func (app *application) updateHandler(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(todoId)

	err := app.todos.Update(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(todoId)

	err := app.todos.Delete(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}
