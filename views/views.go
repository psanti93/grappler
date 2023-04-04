package views

import (
	"fmt"
	"html/template"
	"net/http"
)

type Template struct {
	html *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	err := t.html.Execute(w, data)

	if err != nil {
		fmt.Errorf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
	}

}

func Parse(file string) (*Template, error) {

	tmpl, err := template.ParseFiles(file)

	if err != nil {
		fmt.Errorf("Parsing template %s", err)
		return &Template{}, err
	}

	return &Template{
		html: tmpl,
	}, nil
}
