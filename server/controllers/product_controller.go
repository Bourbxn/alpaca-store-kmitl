package controllers

import (
	"github.com/Bourbxn/alpaca-store-kmitl/config"
	"github.com/Bourbxn/alpaca-store-kmitl/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetPruducts (c *fiber.Ctx) error {
  var products []models.Product
  config.Database.Preload("Options").Find(&products)
  if(len(products)==0){
    return c.Status(404).JSON(fiber.Map{
      "message":"Not found Products",
    })
  }
  return c.Status(200).JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
  idStr := c.Params("id")
  id, err := uuid.Parse(idStr)
  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message":"Invalid ID",
    })
  }
  var product models.Product
  result := config.Database.Preload("Options").Find(&product, id)
  if result.RowsAffected == 0 {
    return c.Status(404).JSON(fiber.Map{
      "message":"Not found User",
    })
  }
  return c.Status(200).JSON(&product)
}

func CreateProduct(c *fiber.Ctx) error {
  product := new(models.Product)
  if err := c.BodyParser(product); err != nil {
    return c.Status(503).SendString(err.Error())
  }
  config.Database.Create(&product)
  return c.Status(201).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
    product := new(models.Product)
    id := c.Params("id")
    if err := c.BodyParser(product); err != nil {
        return c.Status(503).SendString(err.Error())
    }
    config.Database.Where("id = ?", id).Updates(&product)
    return c.Status(200).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
      return c.Status(400).JSON(fiber.Map{
        "message":"Invalid ID",
      })
    }
    var product models.Product
    result := config.Database.Select("Options").Delete(&product, id)
    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.SendStatus(200)
}

