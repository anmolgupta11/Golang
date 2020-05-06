package handlers

import (
	"Go_Project/mysqldb"
	"Go_Project/templates"
	"net/http"

)
// ElementsViewHandler for viewing the elements
func ElementsViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			templates.BasicHTTPRender(w, "Failed to add the element")
			return
		}
		ele := mysqldb.Element{
			Text: r.FormValue("list_object"),
		}
		isAdded := mysqldb.Insert(ele)
		if !isAdded {
			templates.BasicHTTPRender(w, "Failed")
			return
		}
	}

	elements, err := mysqldb.GetAllElements()
	if err != nil {
		templates.BasicHTTPRender(w, "Error reading from database")
		return
	}
	templates.RenderTemplate(w, "templates/elements", &templates.Page{Title: "Elements View", Elements: elements})
}