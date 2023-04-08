package controller

import (
	"fmt"
	"net/http"
)

type User struct {
	Render Renderer
}
type FormData struct {
	Email    string
	Password string
}

func (u *User) NewUser(w http.ResponseWriter, r *http.Request) {
	// run http://localhost:3000/signup?email=user@gmail.com to pre-populate email
	form := &FormData{
		Email: r.FormValue("email"),
	}
	u.Render.Execute(w, form)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "<p> Email is %s</p>", r.PostForm.Get("email"))
	fmt.Fprintf(w, "<p> Password is %s</p>", r.PostForm.Get("password"))

}
