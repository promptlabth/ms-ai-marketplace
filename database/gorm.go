package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/promptlabth/ms-ai-marketplace/app/__mock__/role"
	// agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	// "github.com/promptlabth/ms-ai-marketplace/app/framework"
	// "github.com/promptlabth/ms-ai-marketplace/app/role_framework"
	// "github.com/promptlabth/ms-ai-marketplace/app/user"
)

type GormConnection struct {
	user     string
	password string
	host     string
	name     string
	port     string
}

func NewGormDBWithDefault() *gorm.DB {

	dbConfig := GormConnection{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		password: os.Getenv("DB_PASSWORD"),
		user:     os.Getenv("DB_USER"),
		name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.name)

	return NewGormDB(dsn)
}

func NewGormDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return db
}
