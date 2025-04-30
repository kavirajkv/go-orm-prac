package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kavirajkv/go-orm-prac/src/controller"
)



func MountUser(route fiber.Router){
	usercontroller:=controller.User{}

	route.Post("/createuser",usercontroller.CreateUser)
	route.Get("/getuser",usercontroller.GetUser)
}