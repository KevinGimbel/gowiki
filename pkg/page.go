package gowiki

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
	Slug  string
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("data/"+filename, p.Body, 0600)
}
