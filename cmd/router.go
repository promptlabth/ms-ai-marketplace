package main

import (
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-orch-user-service/app/user"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) {

	userValidation := user.NewAdaptor(db)
	userCore := user.NewCore(db)
	userUsecase := user.NewUsecase(userCore, userValidation)
	userHandler := user.NewHandler(userUsecase)

	router.POST("/user", userHandler.NewUser)
}
