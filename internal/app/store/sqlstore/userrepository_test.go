package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	sqlstore2 "http-rest-api/internal/app/store/sqlstore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore2.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore2.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore2.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore2.New(db)
	email := "user@example.org"
	_, err := s.User().FindBeEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindBeEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
