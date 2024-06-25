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

	runners, err := getAllRunners()
	if err != nil {
		log.Println(err)
	}

	indexTemplate.Execute(w, runners)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	viewTemplate := template.Must(template.ParseFiles("./templates/view.html"))

	runnerId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(runnerId)

	runner, err := getOneRunner(id)
	if err != nil {
		log.Println(err)
	}

	viewTemplate.Execute(w, runner)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	addTemplate := template.Must(template.ParseFiles("./templates/add.html"))
	addTemplate.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		runnerAge := r.FormValue("age")
		age, _ := strconv.Atoi(runnerAge)
		country := r.FormValue("country")
		season_best := r.FormValue("season_best")
		personal_best := r.FormValue("personal_best")

		if name == "" || age == 0 || country == ""  {
			http.Redirect(w, r, "/add", http.StatusMovedPermanently)
		}

		_, err := db.Exec("INSERT INTO runners (name, age, country, season_best, personal_best) VALUES (?, ?, ?, ?, ?)", name, age, country, season_best, personal_best)
		if err != nil {
			log.Println(err)
		}

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	editTemplate := template.Must(template.ParseFiles("./templates/edit.html"))

	runnerId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(runnerId)

	runner, err := getOneRunner(id)
	if err != nil {
		log.Println(err)
	}

	editTemplate.Execute(w, runner)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		runnerAge := r.FormValue("age")
		age, _ := strconv.Atoi(runnerAge)
		country := r.FormValue("country")
		season_best := r.FormValue("season_best")
		personal_best := r.FormValue("personal_best")

		if name == "" || age == 0 || country == ""  {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		_, err := db.Exec("UPDATE runners SET name=?, age=?, country=?, season_best=?, personal_best=? where id=?", name, age, country, season_best, personal_best, id)
		if err != nil {
			log.Println(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := db.Exec("DELETE FROM runners WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}
