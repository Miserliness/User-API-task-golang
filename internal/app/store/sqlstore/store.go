package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/miserliness/users_api_task/internal/app/store"
)

type Store struct {
	db *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			s,
		}
	}

	return s.userRepository
}