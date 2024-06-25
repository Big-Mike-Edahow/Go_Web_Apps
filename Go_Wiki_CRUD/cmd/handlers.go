/* handlers.go */

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("./templates/index.html"))
	var files []string
	
	fileInfo, err := os.ReadDir("./data/")
	if err != nil {
		log.Println(err)
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	indexTemplate.Execute(w, files)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	addTemplate := template.Must(template.ParseFiles("./templates/add.html"))
	addTemplate.Execute(w, nil)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveFileHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	title := r.FormValue("title")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func deleteHandler (w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	dataDir := "./data/"
	file := dataDir + title
	err := os.Remove(file) 
    if err != nil { 
        log.Println(err) 
    } 

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}