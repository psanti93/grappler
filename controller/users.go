package controller

import (
	"net/http"
)

type User struct {
	New Renderer
}

func (u *User) NewUser(w http.ResponseWriter, r *http.Request) {
	u.New.Execute(w, nil)
}
