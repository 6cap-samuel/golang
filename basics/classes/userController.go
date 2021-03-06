package classes

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// constructor
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
