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

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore2.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore2.New(db)
	u1 := model.TestUser(t)

	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore2.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore2.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindBeEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindBeEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
