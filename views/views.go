package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

type Template struct {
	html *template.Template
}

func Must(t *Template, err error) *Template {
	if err != nil {
		fmt.Errorf("Template failed %v", err)
		panic(err)
	}
	return t

}

// Template reflects the Execute function of the Renderer interface
func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.html.Execute(w, data)

	if err != nil {
		fmt.Errorf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
	}

}

func Parse(fs fs.FS, patterns ...string) (*Template, error) {

	tmpl, err := template.ParseFS(fs, patterns...)

	if err != nil {
		fmt.Errorf("Parsing template %s", err)
		return &Template{}, err
	}

	return &Template{
		html: tmpl,
	}, nil
}
