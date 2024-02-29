package main

import (
	"github.com/CRSylar/hattory/cmd"
	"github.com/CRSylar/hattory/handlers"
	"github.com/CRSylar/hattory/view/pages"
	"github.com/a-h/templ"
)

func main() {

	m := handlers.CreateServeMux()

	handlers.Get("/", pages.Layout, []templ.Component{pages.Page(nil)})

	s := cmd.NewServer(m)

	cmd.StartServer(s)
}
