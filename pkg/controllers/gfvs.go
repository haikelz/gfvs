package controllers

import (
	"gfvs/pkg/entities"
	"gfvs/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type GfvsController struct {
	gfvsService *services.GfvsService
}

func NewGfvsController(gfvsService *services.GfvsService) *GfvsController {
	return &GfvsController{gfvsService: gfvsService}
}

func HomeHandler(c *fiber.Ctx) error {
	return c.JSON(entities.Gfvs{
		StatusCode: 200,
		Message:    "success",
	})
}
