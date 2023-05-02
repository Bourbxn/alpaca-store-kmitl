package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {
  app := fiber.New()
  app.Use(cors.New(cors.Config{
		AllowOrigins: "*", 	
    AllowHeaders: "Origin, Content-Type, Accept",
	}))

  app.Get("/", func (c *fiber.Ctx)  error{
    return c.SendString("Hello world")
  })

  app.Listen(":8000")
}
