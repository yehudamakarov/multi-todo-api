package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"multi-todo-api/application/cqhandlers"
	"multi-todo-api/infrastructure/outgoing"
	"multi-todo-api/infrastructure/persistence"
	"multi-todo-api/ui/apicontrollers"
)

func main() {
	per := persistence.NewPersistence()
	authOutgoing := outgoing.NewAuthOutgoing()
	authCqHandler := cqhandlers.NewAuthCqHandler(per, authOutgoing)
	authController := apicontrollers.NewAuthController(authCqHandler)

	r := mux.NewRouter()
	r.HandleFunc("/login", authController.Login)
	r.HandleFunc("/receiveAuthCode", authController.ReceiveAuthCode)
	log.Fatal(http.ListenAndServe(":8080", r))
}
