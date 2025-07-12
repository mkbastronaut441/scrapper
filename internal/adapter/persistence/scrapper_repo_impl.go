package persistence

import (
	"context"
	"log"
	"scrapper/internal/adapter/ports"
	"scrapper/internal/core"

	"github.com/jackc/pgx/v5"
)

type LogsRepoImpl struct {
	db *pgx.Conn
}

func NewLogsRepo(db *pgx.Conn) ports.LogsRepo {

	return &LogsRepoImpl{
		db: db,
	}
}

func (l *LogsRepoImpl) InsertLogs(ctx context.Context, logs core.Logs) (string, error) {
	query := `INSERT INTO logs (level,service,message,metadata) returning event_id;`

	var eventId string
	_, err := l.db.Exec(ctx, query, logs.Level, logs.Service, logs.Message, logs.MetaData)
	if err != nil {
		log.Printf("Failed to insert logs: %v", err)
		return "", err
	}
	return eventId, nil
}
