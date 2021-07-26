package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Sheerlore/gowiki_parctice/wiki"
)

func renderTemplete(res http.ResponseWriter, tmpl string, p *wiki.Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(res, p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	p, err := wiki.LoadPage(title)
	if err != nil {
		http.Redirect(res, req, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplete(res, "../component/view", p)
}

func editHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]
	p, err := wiki.LoadPage(title)
	if err != nil {
		p = &wiki.Page{Title: title}
	}
	renderTemplete(res, "../component/edit", p)
}

func saveHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/save/"):]
	body := req.FormValue("body")
	p := &wiki.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(res, req, "/view/"+title, http.StatusFound)
}

func Run() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
