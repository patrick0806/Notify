package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/patrick0806/orc-notify/internal/entities"
	"github.com/redis/go-redis/v9"
)

func ConsumeNotify(ctx *context.Context, msg *redis.Message) error {
	var notify entities.Notify
	if err := json.Unmarshal([]byte(msg.Payload), &notify); err != nil {
		return err
	}

	slog.Info(
		fmt.Sprintf("Notificação: %s, processada com sucesso", notify.Identity),
		slog.String("operation", "processNotify"),
	)
	return nil
}
