package api

import (
	"fmt"
	"my_first/src/features/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func UserRegister(c *fiber.Ctx)error{
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	  }
	db.AutoMigrate(&UserModel{})
	
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	hash,err := utils.HashPassword(password)

	if err != nil {
		return c.SendString(fmt.Sprintf("error : %s", err))
	}

	db.Create(&UserModel{
		Email: email,
		Password: hash,
		Username: username,
	})

	response := UserResponse {
		Email: email,
		Username: username,
	}

	return c.JSON(response)
}