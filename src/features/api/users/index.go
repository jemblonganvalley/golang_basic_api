package api

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App){
	api := app.Group("/api")
	
	// user
	user := api.Group("/user")
	user.Get("/read", UserRead)
	user.Delete("/delete/:id", UserDelete)
	user.Post("/register", UserRegister)
}