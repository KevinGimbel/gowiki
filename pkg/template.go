package gowiki

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Templates struct{}

// var templates = template.Must(template.ParseFiles("tmpl/baseof.html", "tmpl/view.html", "tmpl/edit.html"))

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("data/" + filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func GetTemplate(tmpl string) (*template.Template, error) {
	return template.ParseFiles("tmpl/baseof.html", "tmpl/"+tmpl+".html")

}

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	tmp, err := GetTemplate(tmpl)
	err = tmp.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
