package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrick0806/orc-notify/internal/entities"
	"github.com/patrick0806/orc-notify/internal/usecases"
)

// @Summary      Receive Notify
// @Description  Receive a notification and save on redis queue
// @Tags         Notify
// @Accept       json
// @Produce      json
// @Param notify body entities.Notify true "Notify"
// @Success      201
// @Failure      400
// @Failure      500
// @Router       /notifies [post]
func ReceiveNotify(c *fiber.Ctx) error {
	c.Locals("operation", "ReceiveNotify")
	var notify entities.Notify
	if err := c.BodyParser(&notify); err != nil {
		errorReturn := entities.Error{
			Code:    400,
			Message: err.Error(),
		}
		return c.Status(400).JSON(errorReturn)
	}

	if err := usecases.SaveNotify(c, notify); err != nil {
		errorReturn := entities.Error{
			Code:    500,
			Message: err.Error(),
		}
		return c.Status(500).JSON(errorReturn)
	}

	return c.SendStatus(201)
}
