/* handlers.go */

package main

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("./templates/index.html")
	prof := Prof{}
	profArrays := []Prof{}

	row, err := conn.Query("SELECT * FROM profs")
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name, age, employ, created string
		err = row.Scan(&id, &name, &age, &employ, &created)
		if err != nil {
			panic(err.Error())
		}

		prof.Id = id
		prof.Name = name
		prof.Age = age
		prof.Employ = employ
		prof.Created = created

		profArrays = append(profArrays, prof)
	}
	indexTemplate.Execute(w, profArrays)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	viewTemplate, _ := template.ParseFiles("./templates/view.html")
	id := r.URL.Query().Get("id")
	prof := Prof{}

	row, err := conn.Query("SELECT id, name, age, employ, created FROM profs where id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name, age, employ, created string
		err = row.Scan(&id, &name, &age, &employ, &created)
		if err != nil {
			panic(err.Error())
		}

		prof.Id = id
		prof.Name = name
		prof.Age = age
		prof.Employ = employ
		prof.Created = created
	}
	viewTemplate.Execute(w, prof)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	addTemplate := template.Must(template.ParseFiles("./templates/add.html"))
	addTemplate.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		employ := r.FormValue("employ")

		if name == "" || age == "" || employ == "" {
			http.Redirect(w, r, "/add", http.StatusMovedPermanently)
		}

		_, err := conn.Exec("INSERT INTO profs (name, age, employ) VALUES (?, ?, ?)", name, age, employ)
		if err != nil {
			panic(err.Error())
		}

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	editTemplate := template.Must(template.ParseFiles("./templates/edit.html"))
	id := r.URL.Query().Get("id")
	prof := Prof{}

	row, err := conn.Query("SELECT id, name, age, employ FROM profs WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name, age, employ string
		err = row.Scan(&id, &name, &age, &employ)
		if err != nil {
			panic(err.Error())
		}

		prof.Id = id
		prof.Name = name
		prof.Age = age
		prof.Employ = employ
	}
	editTemplate.Execute(w, prof)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		age := r.FormValue("age")
		employ := r.FormValue("employ")

		if name == "" || age == "" || employ == "" {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

	
		_, err := conn.Exec("UPDATE profs set name=?, age=?, employ=? where id=?", name, age, employ, id)
		if err != nil {
			panic(err.Error())
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := conn.Exec("DELETE FROM profs WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}
