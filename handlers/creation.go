package handlers

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
)


var m *http.ServeMux

type TempleComponentProps struct {
	Component []templ.Component
	Request *http.Request
}

func CreateServeMux() (http.Handler) {
	// ...
	m = http.NewServeMux()
	
	m.HandleFunc("GET /lib/htmx.js", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			w.Header().Set("Content-Type", "application/gzip")
			http.ServeFile(w, r, "lib/htmx/htmx.min.js.gz")
			return
		}
		w.Header().Set("Content-Type", "application/javascript")
		http.ServeFile(w, r, "lib/htmx/htmx.min.js")
	})

	m.HandleFunc("GET /lib/_hyperscript.js", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			w.Header().Set("Content-Type", "application/gzip")
			http.ServeFile(w, r, "lib/_hyperscript/_hyperscript.min.js.gz")
			return
		}
		w.Header().Set("Content-Type", "application/javascript")
		http.ServeFile(w, r, "lib/_hyperscript/_hyperscript.min.js")
	})
	
	m.HandleFunc("GET /view/public/styles.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "view/public/styles.css")
	})

	return m
}

func Get(path string, tc func(props *TempleComponentProps) templ.Component, c []templ.Component) {
	m.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {
		tc(&TempleComponentProps{c, r}).Render(r.Context(), w)
	}) 
}

func Post(path string, tc func(props *TempleComponentProps) templ.Component, c []templ.Component) {
	m.HandleFunc("POST "+path, func(w http.ResponseWriter, r *http.Request) {
		tc(&TempleComponentProps{c, r}).Render(r.Context(), w)
	})
}

func Delete(path string, tc func(props *TempleComponentProps) templ.Component, c []templ.Component) {
	m.HandleFunc("DELETE "+path, func(w http.ResponseWriter, r *http.Request) {
		tc(&TempleComponentProps{c, r}).Render(r.Context(), w)
	})
}

func Put(path string, tc func(props *TempleComponentProps) templ.Component, c []templ.Component) {
	m.HandleFunc("PUT "+path, func(w http.ResponseWriter, r *http.Request) {
		tc(&TempleComponentProps{c, r}).Render(r.Context(), w)
	})
}