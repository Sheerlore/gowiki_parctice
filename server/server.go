package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Sheerlore/gowiki_parctice/wiki"
)

func renderTemplete(res http.ResponseWriter, tmpl string, p *wiki.Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(res, p)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	p, _ := wiki.LoadPage(title)
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

}

func Run() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
