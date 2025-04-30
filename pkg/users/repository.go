package users

import (
	"github.com/kavirajkv/go-orm-prac/common/schemas"
	"github.com/kavirajkv/go-orm-prac/pkg/models"
)

type Repository interface {
	CreateUser(user *schemas.CreateUser) (*models.User, error)
	GetUserbyEmail(email string) (*models.User, error)
	UpdateUser(user *schemas.UpdateUser) error
}
