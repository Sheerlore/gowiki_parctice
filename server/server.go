package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sheerlore/gowiki_parctice/wiki"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	p, _ := wiki.LoadPage(title)
	fmt.Fprintf(res, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func Run() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
