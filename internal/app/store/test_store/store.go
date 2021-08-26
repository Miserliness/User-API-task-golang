package teststore

import (
	"github.com/miserliness/users_api_task/internal/app/model"
	"github.com/miserliness/users_api_task/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			s,
			make(map[string]*model.User),
		}
	}

	return s.userRepository
}