package model_test

import (
	"rest-api/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValud bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValud: true,
		},
		{
			name: "with encrypt password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedPassword"
				return u
			},
			isValud: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValud: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValud: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValud: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "11111"
				return u
			},
			isValud: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValud {
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
