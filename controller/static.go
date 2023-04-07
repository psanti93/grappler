package controller

import (
	"html/template"
	"net/http"

	"github.com/psanti93/grappler/views"
)

func StaticController(t *views.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}

func FAQ(t *views.Template) http.HandlerFunc {

	question := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is this site used for?",
			Answer:   "It's used to help improve your Jiu Jitsu",
		},
		{
			Question: "Who will review over my skill sets?",
			Answer:   "Your coach",
		},
		{
			Question: "Will others be able to see my skill set?",
			Answer:   "Nope, it stays between you and your coach",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, question)
	}
}
