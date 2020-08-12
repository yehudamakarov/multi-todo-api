package cqhandlers

import "multi-todo-api/application/interfaces"

type authCqHandler struct {
	persistence  interfaces.Persistence
	authOutgoing interfaces.AuthOutgoing
}

type AuthCqHandler interface {
	Login()
	Receive()
}

func NewAuthCqHandler(persistence interfaces.Persistence, outgoing interfaces.AuthOutgoing) AuthCqHandler {
	return &authCqHandler{persistence: persistence, authOutgoing: outgoing}
}

func (a *authCqHandler) Login() {
	panic("implement me")
}

func (a *authCqHandler) Receive() {
	panic("implement me")
}
