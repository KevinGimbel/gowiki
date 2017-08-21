package main

import (
	gowiki "github.com/kevingimbel/gowiki/pkg"
	"net/http"
)

var s = &gowiki.Server{}

func main() {
	http.HandleFunc("/", s.RedirectHome)
	http.HandleFunc("/view/", s.MakeHandler(s.ViewHandler))
	http.HandleFunc("/edit/", s.MakeHandler(s.EditHandler))
	http.HandleFunc("/save/", s.MakeHandler(s.SaveHandler))

	styleHandler := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", styleHandler))

	http.ListenAndServe(":8080", nil)
}
