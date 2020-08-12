package apicontrollers

import (
	"net/http"

	"multi-todo-api/application/cqhandlers"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	ReceiveAuthCode(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authCqHandler cqhandlers.AuthCqHandler
}

func NewAuthController(authCqHandler cqhandlers.AuthCqHandler) AuthController {
	return &authController{authCqHandler: authCqHandler}
}

func (ac *authController) Login(w http.ResponseWriter, r *http.Request) {
	url := "url here"
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (ac *authController) ReceiveAuthCode(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
