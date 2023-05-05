package routes

import (
	"github.com/Bourbxn/alpaca-store-kmitl/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
  router.Get("/products", controllers.GetPruducts)
  router.Get("/product/:id", controllers.GetProduct)
  router.Post("/product", controllers.CreateProduct)
  router.Put("/product/:id", controllers.UpdateProduct)
  router.Delete("/product/:id", controllers.DeleteProduct)
}
