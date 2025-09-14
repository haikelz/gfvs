package routes

import (
	"gfvs/pkg/configs"
	"gfvs/pkg/controllers"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HomeRoute(api fiber.Router) {
	api.Get("/", controllers.HomeHandler)
}

func PrometheusRoute(api fiber.Router) {
	api.Use("/metrics", adaptor.HTTPHandler(controllers.PrometheusHandler()))
}

func SwaggerRoute(api fiber.Router) {
	api.Use(swagger.New(configs.SwgCfg))
}
