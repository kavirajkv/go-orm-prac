package migrations

import (
	"fmt"

	"github.com/kavirajkv/go-orm-prac/common/db"
	"github.com/kavirajkv/go-orm-prac/pkg/models"
)

func Migrate() {
	database := db.InitDB()
	if database == nil {
		panic("Database connection failed")
	}

	err := database.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated successfully")
}
