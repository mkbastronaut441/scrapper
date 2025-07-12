package usecase

import (
	"context"
	"scrapper/internal/adapter/ports"
	"scrapper/internal/core"
)

type LogsUsecaseImpl struct {
	logsRepo ports.LogsRepo
}

func NewLogsUsecaseImpl(logsRepo ports.LogsRepo) core.LogsUsecase {
	return &LogsUsecaseImpl{logsRepo: logsRepo}
}

func (l *LogsUsecaseImpl) StoreLogs(ctx context.Context, logs core.Logs) (string, error) {

	eventId, err := l.logsRepo.InsertLogs(ctx, logs)
	if err != nil {
		return "", err
	}
	return eventId, err
}
