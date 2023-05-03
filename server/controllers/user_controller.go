package controllers

import (
	// "github.com/Bourbxn/alpaca-store-kmitl/config"
	// "github.com/Bourbxn/alpaca-store-kmitl/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
    // var users []models.User
    // config.Database.Find(&users)
    return c.SendString("get users route completed")
}
