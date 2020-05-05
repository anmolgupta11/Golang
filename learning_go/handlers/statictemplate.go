package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func StaticTemplates() {
	http.HandleFunc("/", serveTemplate)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/view/", http.StripPrefix("/view/", fs))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path)) //filepath.Clean() is used to sanitise the URL path before using it, because the URL path is untrusted user input.

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)

	// return a 404 error if template does not exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// return 404 error if request is from a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}
}
