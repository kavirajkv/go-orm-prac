package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kavirajkv/go-orm-prac/common/config"
	"github.com/kavirajkv/go-orm-prac/src/router"

	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RunServer() {

	app := fiber.New(fiber.Config{
		AppName: "Go ORM Practice",
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} - ${latency} ${method} ${path} ${ip}\n"}))

	router.Routes(app)

	port := config.PORT
	app.Listen(fmt.Sprintf(":%s", port))
	fmt.Printf("Server is running on port %s\n", port)
}
