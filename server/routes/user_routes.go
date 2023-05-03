package routes

import (
	"github.com/Bourbxn/alpaca-store-kmitl/controllers"
	"github.com/gofiber/fiber/v2"
)


func UsersRoutes(router fiber.Router) {
  router.Get("users", controllers.GetUsers)
}
