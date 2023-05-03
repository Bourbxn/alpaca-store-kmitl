package controllers

import (
	"github.com/Bourbxn/alpaca-store-kmitl/config"
	"github.com/Bourbxn/alpaca-store-kmitl/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
  var users []models.User
  config.Database.Find(&users)
  if(len(users)==0){
    return c.Status(404).JSON(fiber.Map{
      "message":"Not found Users",
    })
  }
  return c.Status(200).JSON(users)
}
