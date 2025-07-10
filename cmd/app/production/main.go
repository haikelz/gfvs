package app

import (
	"net/http"

	"gfvs/pkg/configs"
	"gfvs/pkg/server"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// @title Go Fiber Vercel Starter
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// Handle deployment to Vercel Serverless
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	server := server.New()
	server.RegisterFiberRoutes()
	server.Use(swagger.New(configs.SwgCfg))

	adaptor.FiberApp(server.App).ServeHTTP(w, r)
}
