package outgoing

import "multi-todo-api/application/interfaces"

type authOutgoing struct {
}

func NewAuthOutgoing() interfaces.AuthOutgoing {
	return &authOutgoing{}
}
