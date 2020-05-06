package handlers

import (
	"Go_Project/templates"
	"net/http"
)

// HomeViewHandler to check for Home Page
func HomeViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	t := &templates.Page{Title: "Home"}
	templates.RenderTemplate(w, "templates/home", t)
}
