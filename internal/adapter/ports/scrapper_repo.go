package ports

import (
	"context"
	"scrapper/internal/core"
)

type LogsRepo interface {
	InsertLogs(ctx context.Context, logs core.Logs) (string, error)
}
