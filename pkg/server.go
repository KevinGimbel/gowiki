package gowiki

import (
	"net/http"
)

type Server struct{}

func (s *Server) ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	RenderTemplate(w, "view", p)
}

func (s *Server) EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	RenderTemplate(w, "edit", p)
}

func (s *Server) SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func (s *Server) FrontpageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := LoadPage("frontpage")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	RenderTemplate(w, "static", p)
}

func (s *Server) RedirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/frontpage", http.StatusFound)
}

func (s *Server) MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
