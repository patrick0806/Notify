package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/patrick0806/orc-notify/internal"
	"github.com/patrick0806/orc-notify/internal/config"
	"github.com/patrick0806/orc-notify/internal/middleware"

	_ "github.com/patrick0806/orc-notify/docs"
)

// @title           Orc Notify
// @version         1.0
// @description     This is a orc how recive a notify and send this to a queue and process.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	config.Logger()
	app := fiber.New()
	app.Use(middleware.LoggerMiddleware)
	app.Get("/swagger/*", swagger.HandlerDefault)
	internal.Router(app)
	app.Listen(":3000")
	slog.Info("Server running on port 3000")
}
