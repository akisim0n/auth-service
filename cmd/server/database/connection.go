package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type Config struct {
	DSN string
}

func newConfig(config string) *Config {
	return &Config{DSN: config}
}

func Connect(ctx context.Context) (*pgxpool.Pool, error) {

	config := newConfig(fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		getEnv("PG_DB_HOST"),
		getEnv("PG_PORT_IN"),
		getEnv("PG_DB_NAME"),
		getEnv("PG_USER"),
		getEnv("PG_PASSWORD")))

	dbPool, err := pgxpool.New(ctx, config.DSN)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error connecting to database: %v", err))
	}

	return dbPool, nil
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}
