package main

import (
	"github.com/CRSylar/red_tetris/cmd"
	"github.com/CRSylar/red_tetris/handlers"
	"github.com/CRSylar/red_tetris/view/pages"
	"github.com/a-h/templ"
)

func main() {

	m := handlers.CreateServeMux()

	handlers.Get("/", pages.Layout, []templ.Component{pages.Page(nil)})

	s := cmd.NewServer(m)

	cmd.StartServer(s)
}