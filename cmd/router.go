package main

import (
	"database/sql"

	"github.com/promptlabth/ms-orch-user-service/app"
	"github.com/promptlabth/ms-orch-user-service/app/user"
)

func NewRouter(router *app.RouterGin, db *sql.DB) {

	userValidation := user.NewAdaptor(db)
	userCore := user.NewCore(db)
	userUsecase := user.NewUsecase(userCore, userValidation)
	userHandler := user.NewHandler(userUsecase)

	router.POST("/user", userHandler.NewUser)
}
