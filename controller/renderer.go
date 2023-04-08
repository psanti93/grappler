package controller

import "net/http"

type Renderer interface {
	Execute(w http.ResponseWriter, data interface{})
}
