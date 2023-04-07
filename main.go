package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/psanti93/grappler/controller"
	"github.com/psanti93/grappler/templates"
	"github.com/psanti93/grappler/views"
)

func main() {
	r := chi.NewRouter()

	tmpl := views.Must(views.Parse(templates.FS, "home.gohtml"))
	r.Get("/", controller.StaticController(tmpl))

	tmpl = views.Must(views.Parse(templates.FS, "signup.gohtml"))

	r.Get("/signup", controller.StaticController(tmpl))

	tmpl = views.Must(views.Parse(templates.FS, "faq.gohtml"))
	r.Get("/faq", controller.FAQ(tmpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})
	http.ListenAndServe(":3000", r)
}
