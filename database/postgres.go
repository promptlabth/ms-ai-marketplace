package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgres() *sql.DB {
	dbConfig := GormConnection{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		password: os.Getenv("DB_PASSWORD"),
		user:     os.Getenv("DB_USER"),
		name:     os.Getenv("DB_NAME"),
	}
	dsn := fmt.Sprintf("postgres://%s:%s@(%s:%s)/%s?sslmode=disable", dbConfig.user, dbConfig.password, dbConfig.host, dbConfig.port, dbConfig.name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Panic(err)
	}
	return db
}
