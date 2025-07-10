package main

import (
	"gfvs/pkg/configs"
	"gfvs/pkg/server"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title Go Fiber Vercel Starter
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	server := server.New()
	server.RegisterFiberRoutes()

	server.Use("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	server.Use(swagger.New(configs.SwgCfg))

	server.Listen(":4000")
}
