package database

import (
	"fmt"
	"log"

	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormConnection struct {
	user     string
	password string
	host     string
	name     string
	port     string
}

func NewGormDBWithDefault() *gorm.DB {

	// dbConfig := GormConnection{
	// 	host:     os.Getenv("DB_HOST"),
	// 	port:     os.Getenv("DB_PORT"),
	// 	password: os.Getenv("DB_PASSWORD"),
	// 	user:     os.Getenv("DB_USER"),
	// 	name:     os.Getenv("DB_NAME"),
	// }
	dbConfig := GormConnection{
		host:     "localhost",
		port:     "5432",
		password: "1234",
		user:     "prompty",
		name:     "prompty",
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.name)

	return NewGormDB(dsn)
}

func NewGormDB(dns string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db
}
