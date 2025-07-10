package controllers

import (
	"gfvs/pkg/models"
	"gfvs/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type GfvsController struct {
	gfvsService *services.GfvsService
}

func NewGfvsController(gfvsService *services.GfvsService) *GfvsController {
	return &GfvsController{gfvsService: gfvsService}
}

// HomeHandler godoc
// @Summary Get Home
// @Description Get Home
// @Tags Home
// @Accept json
// @Produce json
// @Success 200 {object} models.APIHomeResponse "Get Home Success"
// @Failure 400 {object} models.APIErrorResponse "Bad Request"
// @Failure 500 {object} models.APIErrorResponse "Internal Server Error"
// @Router / [get]
func HomeHandler(c *fiber.Ctx) error {
	return c.JSON(models.APIHomeResponse{
		StatusCode: 200,
		Message:    "success",
	})
}
