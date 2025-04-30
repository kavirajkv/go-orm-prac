package db

import (
	"fmt"
	"github.com/kavirajkv/go-orm-prac/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	db := InitDB()
	return db
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", config.DB_HOST, config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	return db

}
