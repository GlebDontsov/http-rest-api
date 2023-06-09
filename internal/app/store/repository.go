package store

import "http-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindBeEmail(string) (*model.User, error)
}
