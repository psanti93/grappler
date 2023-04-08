package controller

import (
	"net/http"

	"github.com/psanti93/grappler/views"
)

type User struct {
	New *views.Template
}

func (u *User) NewUser(w http.ResponseWriter, r *http.Request) {
	u.New.Execute(w, nil)
}
