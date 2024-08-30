package middleware

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	defer func() {
		path := c.Path()

		if ignorePath(path) {
			return
		}

		duration := time.Since(start)
		status := c.Response().StatusCode()
		method := c.Method()
		operation := c.Locals("operation")
		if operation == nil {
			errorMessage := fmt.Sprintf("Requisição de path %s não tem operação definida", path)
			slog.Error(errorMessage)
			c.Response().SetStatusCode(500)
			return
		}

		hostanme, err := os.Hostname()
		if err != nil {
			hostanme = "unknown"
		}

		if status >= 400 {
			slog.Warn(
				"Request Warn",
				slog.String("operation", operation.(string)),
				slog.String("method", method),
				slog.String("path", path),
				slog.Int("status", status),
				slog.Duration("duration", duration),
				slog.String("error", c.Response().String()),
				slog.String("remote_ip", c.IP()),
				slog.String("hostname", hostanme),
			)
			return
		}

		if status >= 500 {
			slog.Error(
				"Request error",
				slog.String("operation", operation.(string)),
				slog.String("method", method),
				slog.String("path", path),
				slog.Int("status", status),
				slog.Duration("duration", duration),
				slog.String("error", c.Response().String()),
				slog.String("remote_ip", c.IP()),
				slog.String("hostname", hostanme),
			)
			return
		}

		slog.Info(
			"Request successful",
			slog.String("operation", operation.(string)),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status", status),
			slog.Duration("duration", duration),
			slog.String("remote_ip", c.IP()),
			slog.String("hostname", hostanme),
		)
	}()
	return c.Next()
}

func ignorePath(path string) bool {
	paths := []string{"health", "swagger", "png", "json"}
	for _, pattern := range paths {
		if strings.Contains(path, pattern) {
			return true
		}
	}

	return false
}
