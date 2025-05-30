package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DSN string
}

func newConfig(config string) *Config {
	return &Config{DSN: config}
}

func Connect(ctx context.Context) (*pgxpool.Pool, error) {

	if envErr := godotenv.Load(); envErr != nil {
		return nil, errors.New(fmt.Sprintf("Error loading .env file: %v", envErr))
	}

	config := newConfig(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		getEnv("PG_DB_HOST"),
		getEnv("PG_PORT"),
		getEnv("PG_USER"),
		getEnv("PG_PASSWORD"),
		getEnv("PG_DB_NAME")))

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
