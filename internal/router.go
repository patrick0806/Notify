package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrick0806/orc-notify/internal/controller"
)

func Router(app *fiber.App) {
	app.Get("/health", controller.Health)
	app.Post("/notifies", controller.ReceiveNotify)
}
