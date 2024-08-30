package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrick0806/orc-notify/internal/datasource"
)

type HealthResponse struct {
	Server string `json:"server"`
	Redis  string `json:"redis"`
}

// @Summary      Health Check
// @Description  Check orc health
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  HealthResponse
// @Failure      500
// @Router       /health [get]
func Health(c *fiber.Ctx) error {
	ctx := c.Context()
	redisClient := datasource.GetRedisClient()
	_, err := redisClient.Ping(ctx).Result()

	redisStatus := "ok"
	if err != nil {
		redisStatus = "error"
	}

	response := HealthResponse{
		Server: "ok",
		Redis:  redisStatus,
	}

	return c.JSON(response)
}
