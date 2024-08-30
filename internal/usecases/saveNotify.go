package usecases

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/patrick0806/orc-notify/internal/datasource"
	"github.com/patrick0806/orc-notify/internal/entities"
)

func SaveNotify(ctx *fiber.Ctx, notify entities.Notify) error {
	redisClient := datasource.GetRedisClient()

	jsonNotify, err := json.Marshal(notify)
	if err != nil {
		return err
	}

	_, err = redisClient.Publish(ctx.Context(), "notify", string(jsonNotify)).Result()
	if err != nil {
		return err
	}

	return nil
}
