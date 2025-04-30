package users

import (
	"github.com/kavirajkv/go-orm-prac/common/schemas"
	"github.com/kavirajkv/go-orm-prac/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) CreateUser(obj *schemas.CreateUser) (*models.User, error) {
	user := models.User{
		Username: obj.Username,
		Password: obj.Password,
		Email:    obj.Email,
		Phone:    obj.Phone,
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) GetUserbyEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *repo) UpdateUser(obj *schemas.UpdateUser) error {
	user := models.User{
		Username: obj.Username,
		Password: obj.Password,
		Email:    obj.Email,
		Phone:    obj.Phone,
	}
	if err := r.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil

}

func NewRepo(db *gorm.DB) Repository {
	return &repo{DB: db}
}
