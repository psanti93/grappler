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

	tmpl := views.Must(views.Parse(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controller.StaticController(tmpl))

	tmpl = views.Must(views.Parse(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controller.StaticController(tmpl))

	tmpl = views.Must(views.Parse(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controller.FAQ(tmpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})

	var usersC controller.User
	// views returns type Template which reflects the Renderer interface in the controllers package and assigns to Render of type rendere interface
	usersC.Render = views.Must(views.Parse(templates.FS, "signup.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.NewUser)
	r.Post("/signup", usersC.CreateUser)

	http.ListenAndServe(":3000", r)
}
