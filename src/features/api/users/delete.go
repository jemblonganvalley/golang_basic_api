package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func UserDelete(c *fiber.Ctx)error{
	id := c.Params("id")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var user UserModel
	result := db.First(&user, id)
	if result.Error != nil {
		c.SendStatus(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"success" : false,
		})
	}

	db.Delete(&user)
	respose := fmt.Sprintf("Delete user %s", id)
	return c.SendString(respose)
}