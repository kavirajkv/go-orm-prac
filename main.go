package main

import (
	"github.com/kavirajkv/go-orm-prac/common/config"
	"github.com/kavirajkv/go-orm-prac/common/db"
	"github.com/kavirajkv/go-orm-prac/common/migrations"
	"github.com/kavirajkv/go-orm-prac/src/fiber"

)

func main() {
	config.InitConfig()

	db:=db.InitDB()

	//db migration
	migrations.Migrate(db)

	fiber.RunServer()

}