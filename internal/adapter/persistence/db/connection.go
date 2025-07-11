package persistence

import (
	"context"
	"fmt"
	"scrapper/internal/config"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func NewDbConnection(conn *config.DBConfig) *Postgres {
	return &Postgres{
		DBHost:     conn.DBHost,
		DBPort:     conn.DBPort,
		DBName:     conn.DBName,
		DBUser:     conn.DBUser,
		DBPassword: conn.DBPassword,
	}

}

func (p Postgres) Connect() (*pgx.Conn, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.DBHost, p.DBPort, p.DBUser, p.DBPassword, p.DBName,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("connected to db")

	return conn, nil
}
