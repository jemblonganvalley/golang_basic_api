package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func UserRead(c *fiber.Ctx)error{
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	  }

	var users []UserModel
	db.Find(&users)

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID: int(user.ID),
			Email: user.Email,
			Username: user.Username,
		})
	}

	// response := UserResponse{
	// 	Success: true,
	// 	Email: user.Email,
	// 	Username: user.Username,
	// }

	return c.JSON(userResponses)
}