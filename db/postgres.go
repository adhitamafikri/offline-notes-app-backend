package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Postgres() {
	env := os.Getenv("APP_ENV")
	envFile := fmt.Sprintf(".env.%s", env)
	_ = godotenv.Load(envFile)

	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", postgresUser, postgresPassword, postgresDB)

	// conn, err := pgx.Connect(context.Background(), postgresURL)
	dbpool, err := pgxpool.New(context.Background(), postgresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var name string
	var weight int64
	err = dbpool.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
