package migrations

import (
	"fmt"

	"github.com/kavirajkv/go-orm-prac/pkg/models"
	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	if database==nil{
		panic("Database connection failed")
	}

	err:=database.AutoMigrate(&models.User{},&models.Account{},&models.Transaction{})
	if err!=nil{
		panic(err)
	}

	fmt.Println("Database migrated successfully")
}