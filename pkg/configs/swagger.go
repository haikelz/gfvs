package configs

import (
	"github.com/gofiber/contrib/swagger"
)

var SwgCfg = swagger.Config{
	BasePath:    "/",
	Title:       "Go Fiber Vercel Starter API Docs",
	Path:        "docs",
	FileContent: []byte(SwaggerJSON),
	CacheAge:    60,
}
