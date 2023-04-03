package store_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"http-rest-api/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
