package teststore_test

import (
	"github.com/miserliness/users_api_task/internal/app/model"
	"github.com/miserliness/users_api_task/internal/app/store"
	teststore "github.com/miserliness/users_api_task/internal/app/store/test_store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "usertest@example.ru"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	s.User().Create(u)

	u, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_DeleteByID(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	s.User().Create(u)

	u, err := s.User().Find(u.ID)
	assert.NotNil(t, u)

	err = s.User().DeleteByID(u.ID)
	u, err = s.User().Find(u.ID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	assert.Nil(t, u)
}
