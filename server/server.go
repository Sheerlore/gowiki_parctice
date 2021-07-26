package server

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/Sheerlore/gowiki_parctice/wiki"
)

const pagePath = "../component/"

var templetes = template.Must(template.ParseFiles(pagePath+"edit.html", pagePath+"view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplete(res http.ResponseWriter, tmpl string, p *wiki.Page) {
	err := templetes.ExecuteTemplate(res, tmpl+".html", p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle(res http.ResponseWriter, req *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(req.URL.Path)
	if m == nil {
		http.NotFound(res, req)
		return "", errors.New("invaild Page Title")
	}
	return m[2], nil // the title is seconde subexpression.
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	title, err := getTitle(res, req)
	if err != nil {
		return
	}
	p, err := wiki.LoadPage(title)
	if err != nil {
		http.Redirect(res, req, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplete(res, pagePath+"view", p)
}

func editHandler(res http.ResponseWriter, req *http.Request) {
	title, err := getTitle(res, req)
	if err != nil {
		return
	}
	p, err := wiki.LoadPage(title)
	if err != nil {
		p = &wiki.Page{Title: title}
	}
	renderTemplete(res, pagePath+"edit", p)
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
