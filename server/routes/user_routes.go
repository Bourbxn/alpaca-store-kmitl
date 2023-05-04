package routes

import (
	"github.com/Bourbxn/alpaca-store-kmitl/controllers"
	"github.com/gofiber/fiber/v2"
)


func UsersRoutes(router fiber.Router) {
  router.Get("/users", controllers.GetUsers)
  router.Get("/user/:id", controllers.GetUser)
  router.Post("/user", controllers.CreateUser)
  router.Put("/user/:id", controllers.UpdateUser)
  router.Delete("/user/:id", controllers.DeleteUser)
}
