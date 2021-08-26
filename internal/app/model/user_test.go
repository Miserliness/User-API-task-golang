package model_test

import (
	"github.com/miserliness/users_api_task/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct{
		name string
		u func() *model.User
		isValid bool
	}{
		{
			"valid",
			func() *model.User{
				return model.TestUser(t)
			},
			true,
		},
		{
			"empty email",
			func() *model.User{
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			false,
		},
		{
			"invalid email",
			func() *model.User{
				u := model.TestUser(t)
				u.Email = "invalid.email"
				return u
			},
			false,
		},
		{
			"empty password",
			func() *model.User{
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			false,
		},
		{
			"short password",
			func() *model.User{
				u := model.TestUser(t)
				u.Password = "pass"
				return u
			},
			false,
		},
		{
			"with enc pass",
			func() *model.User{
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpaswword"
				return u
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
