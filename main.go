package main

import (
	"github.com/Sheerlore/gowiki_parctice/server"
	"github.com/Sheerlore/gowiki_parctice/wiki"
)

func main() {
	p1 := &wiki.Page{Title: "TestPage", Body: []byte("This is test")}
	p1.Save()

	server.Run()
}
