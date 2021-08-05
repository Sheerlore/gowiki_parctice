package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/Sheerlore/gowiki_parctice/wiki"
)

var templetes = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplete(res http.ResponseWriter, tmpl string, p *wiki.Page) {
	err := templetes.ExecuteTemplate(res, tmpl+".html", p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// func getTitle(res http.ResponseWriter, req *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(req.URL.Path)
// 	if m == nil {
// 		http.NotFound(res, req)
// 		return "", errors.New("invaild Page Title")
// 	}
// 	return m[2], nil // the title is seconde subexpression.
// }

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		m := validPath.FindStringSubmatch(req.URL.Path)
		if m == nil {
			http.NotFound(res, req)
			return
		}
		fn(res, req, m[2])
	}
}

func viewHandler(res http.ResponseWriter, req *http.Request, title string) {
	p, err := wiki.LoadPage(title)
	if err != nil {
		http.Redirect(res, req, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplete(res, "view", p)
}

func editHandler(res http.ResponseWriter, req *http.Request, title string) {
	p, err := wiki.LoadPage(title)
	if err != nil {
		p = &wiki.Page{Title: title}
	}
	renderTemplete(res, "edit", p)
}

func saveHandler(res http.ResponseWriter, req *http.Request, title string) {
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
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
