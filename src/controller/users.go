package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kavirajkv/go-orm-prac/common/db"
	"github.com/kavirajkv/go-orm-prac/common/schemas"
)

type User struct {}

func(u User)CreateUser(c *fiber.Ctx)error{
	var user schemas.CreateUser

	err:=c.BodyParser(&user)

	if err!=nil{
		return c.Status(400).SendString("Invalid request")
	}

	_,er:=db.Userservice.CreateUser(&user)
	if er!=nil{
		return c.Status(500).SendString("Internal server error")
	}
	return c.Status(200).SendString("User created successfully")

}

func(u User)GetUser(c *fiber.Ctx)error{
	var email schemas.Getuser

	err:=c.BodyParser(&email)

	if err!=nil{
		return c.Status(400).SendString("Invalid request")
	}

	user,err:=db.Userservice.GetUserbyEmail(email.Email)

	if err!=nil{
		return c.Status(500).SendString("Internal server error")
	}
	return c.Status(200).JSON(&user)

}