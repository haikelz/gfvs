package routes

import (
	"gfvs/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func HomeRoute(api fiber.Router) {
	api.Get("/", controllers.HomeHandler)
}
