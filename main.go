package main

import (
	"github.com/kavirajkv/go-orm-prac/common/config"
	"github.com/kavirajkv/go-orm-prac/common/db"
	"github.com/kavirajkv/go-orm-prac/common/migrations"
	"github.com/kavirajkv/go-orm-prac/src/fiber"
)

func main() {
	config.InitConfig()

	//initialize db connection
	db.InitServices()

	//db migration
	migrations.Migrate()

	fiber.RunServer()

}