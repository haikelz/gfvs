package main

import (
	"gfvs/pkg/server"
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

	server.Listen(":4000")
}
