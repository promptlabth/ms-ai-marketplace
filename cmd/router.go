package main

import (
	"database/sql"

	"github.com/promptlabth/ms-orch-user-service/app"
	"github.com/promptlabth/ms-orch-user-service/app/user"
)

func NewRouter(router *app.RouterGin, db *sql.DB) {

	userStorage := user.NewAdaptor(db)
	userCore := user.NewCore(db)
	userUsecase := user.NewUsecase(userStorage, userCore)
	userHandler := user.NewHandler(userUsecase)

	router.GET("/", userHandler.NewUser)
}
