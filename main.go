package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/psanti93/grappler/views"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tmplPath)

}

func signUpPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "signup.gohtml")
	executeTemplate(w, tmplPath)
}

func executeTemplate(w http.ResponseWriter, filePath string) {
	tmpl, err := views.Parse(filePath)

	if err != nil {
		fmt.Errorf("There was an issue parsing the file: %v", err)
		http.Error(w, "Files could not be parsed", http.StatusInternalServerError)
	}

	tmpl.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homePage)
	r.Get("/signup", signUpPage)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})
	http.ListenAndServe(":3000", r)
}
