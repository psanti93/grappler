package controller

import (
	"net/http"
)

type User struct {
	Render Renderer
}

func (u *User) NewUser(w http.ResponseWriter, r *http.Request) {
	u.Render.Execute(w, nil)
}
