package schemas

import (
	"github.com/kavirajkv/go-orm-prac/pkg/models"
)

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateUser struct {
	Username string `json:"username" validate:"omitempty,min=3,max=20"`
	Password string `json:"password" validate:"omitempty,min=8,max=20"`
	Email    string `json:"email" validate:"omitempty,email"`
	Phone    string `json:"phone" validate:"omitempty"`
}

type UserResponse struct {
	ID        int64            `json:"id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Phone     string           `json:"phone"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
	Accounts  []models.Account `json:"accounts"`
}


type Getuser struct {
	Email string `json:"email"`
}
