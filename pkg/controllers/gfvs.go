package controllers

import (
	"gfvs/pkg/models"
	"gfvs/pkg/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		StatusCode: http.StatusOK,
		Message:    "success get home",
	})

}

// PrometheusHandler godoc
// @Summary Get Metrics
// @Description Get Metrics
// @Tags Metrics
// @Accept json
// @Produce json
// @Failure 400 {object} models.APIErrorResponse "Bad Request"
// @Failure 500 {object} models.APIErrorResponse "Internal Server Error"
// @Router /metrics [get]
func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}
