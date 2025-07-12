package core

import (
	"context"

	"github.com/google/uuid"
)

type Logs struct {
	Id       int                    `json:"id"`
	EventId  uuid.UUID              `json:"eventId"`
	Level    string                 `json:"level"`
	Service  string                 `json:"service"`
	Message  string                 `json:"message"`
	MetaData map[string]interface{} `json:"metadata"`
}

type LogsUsecase interface {
	StoreLogs(ctx context.Context, logs Logs) (string, error)
}
