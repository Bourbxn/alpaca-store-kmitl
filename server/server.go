package main

import (
	"log"
	"github.com/Bourbxn/alpaca-store-kmitl/config"
	"github.com/Bourbxn/alpaca-store-kmitl/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {
  app := fiber.New()

  api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: "*", 	
    AllowHeaders: "Origin, Content-Type, Accept",
	}))

  config.Connect()
  
  routes.UsersRoutes(api)

  log.Fatal(app.Listen(":5500"))
}
