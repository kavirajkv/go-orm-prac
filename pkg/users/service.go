package users

import (
	"github.com/kavirajkv/go-orm-prac/common/schemas"
	"github.com/kavirajkv/go-orm-prac/pkg/models"
)

type usersvc struct {
	repo Repository
}

type Service interface {
	CreateUser(user *schemas.CreateUser) (*models.User, error)
	GetUserbyEmail(email string) (*models.User, error)
	UpdateUser(user *schemas.UpdateUser) error
}

func (s *usersvc) CreateUser(user *schemas.CreateUser) (*models.User, error) {
	return s.repo.CreateUser(user)
}

func (s *usersvc) GetUserbyEmail(email string) (*models.User, error) {
	return s.repo.GetUserbyEmail(email)
}

func (s *usersvc) UpdateUser(user *schemas.UpdateUser) error {
	return s.repo.UpdateUser(user)
}

////////////////
//service to interact with the repository

func Newservice(r Repository) Service {
	return &usersvc{repo: r}
}
