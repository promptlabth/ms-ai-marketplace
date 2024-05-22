package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// agentdetail "github.com/promptlabth/ms-orch-user-service/app/agent_detail"
	// "github.com/promptlabth/ms-orch-user-service/app/framework"
	// "github.com/promptlabth/ms-orch-user-service/app/role_framework"
	// "github.com/promptlabth/ms-orch-user-service/app/user"
	// "github.com/promptlabth/ms-orch-user-service/app/__mock__/role"
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
		password: "myPasswordAtPromptLabAI",
		user:     "promptlabai",
		name:     "promptlabai-db",
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.password, dbConfig.name)

	return NewGormDB(dsn)
}

func NewGormDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	// if err := db.AutoMigrate(
	// 	&user.UserEntity{},
	// 	&agentdetail.AgentDetailEntity{},
	// 	&framework.FrameworkEntity{},
	// 	&roleframework.RoleFrameworkEntity{},
	// 	&role.RoleEntity{}); err != nil { // Add RoleEntity to the migrations
	// 	log.Fatalf("Failed to run migrations: %v", err)
	// }

	return db
}
