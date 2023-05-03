package controllers

import (
	"github.com/Bourbxn/alpaca-store-kmitl/config"
	"github.com/Bourbxn/alpaca-store-kmitl/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func GetUser(c *fiber.Ctx) error {
  idStr := c.Params("id")
  id, err := uuid.Parse(idStr)
  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message":"Invalid ID",
    })
  }
  var user models.User
  result := config.Database.Find(&user, id)
  if result.RowsAffected == 0 {
    return c.Status(404).JSON(fiber.Map{
      "message":"Not found User",
    })

  }
  return c.Status(200).JSON(&user)
}
