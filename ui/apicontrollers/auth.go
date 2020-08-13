package apicontrollers

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"multi-todo-api/application/cqhandlers"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	ReceiveAuthCode(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authCqHandler cqhandlers.AuthCqHandler
}

// todo - randomize
var state string = "randomized string"
var config = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	Endpoint:     google.Endpoint,
	RedirectURL:  "",
	Scopes:       nil,
}

func NewAuthController(authCqHandler cqhandlers.AuthCqHandler) AuthController {
	return &authController{authCqHandler: authCqHandler}
}

func (ac *authController) Login(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (ac *authController) ReceiveAuthCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	state := vars["state"]
}
