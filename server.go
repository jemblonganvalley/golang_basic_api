package main

import (
	api "my_first/src/features/api/users"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	api.UserRoutes(app)

	app.Listen(":3000")

}