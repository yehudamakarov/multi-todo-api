package persistence

import "multi-todo-api/application/interfaces"

type mongoRepository struct {

}

func NewPersistence() interfaces.Persistence {
	return &mongoRepository{}
}

func (m *mongoRepository) StoreUser() {
	panic("implement me")
}



