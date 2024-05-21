package main

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-orch-user-service/app/agent_detail"
	"github.com/promptlabth/ms-orch-user-service/app/framework"
	// "github.com/promptlabth/ms-orch-user-service/app/role"
	"github.com/promptlabth/ms-orch-user-service/app/upload"
	"github.com/promptlabth/ms-orch-user-service/app/user"
	"github.com/promptlabth/ms-orch-user-service/app/__mock__/role"
	"gorm.io/gorm"
)

func AgentDetailRouter(router *gin.Engine, db *gorm.DB) {
	agentDetailValidation := agentdetail.NewAdaptor(db)
	agentDetailCore := agentdetail.NewCore(db)
	agentDetailUsecase := agentdetail.NewUsecase(agentDetailCore, agentDetailValidation)
	agentDetailHandler := agentdetail.NewHandler(agentDetailUsecase)

	router.POST("/creator/agent_detail", agentDetailHandler.NewAgentDetail)
	router.GET("/creator/:id", agentDetailHandler.GetAgentDetails)
}

func FrameworkRouter(router *gin.Engine, db *gorm.DB) {
	frameworkValidation := framework.NewAdaptor(db)
	frameworkCore := framework.NewCore(db)
	frameworkUsecase := framework.NewUsecase(frameworkCore, frameworkValidation)
	frameworkHandler := framework.NewHandler(frameworkUsecase)

	router.POST("/creator/framework", frameworkHandler.NewFramework)
	router.GET("/creator/frameworks", frameworkHandler.ListFrameworks)
}

func RoleRouter(router *gin.Engine, db *gorm.DB) {

	roleValidation := role.NewAdaptor(db)
	roleCore := role.NewCore(db)
	roleUsecase := role.NewUsecase(roleCore, roleValidation)
	roleHandler := role.NewHandler(roleUsecase)

	router.POST("/creator/role", roleHandler.NewRole)
}

func UserRouter(router *gin.Engine, db *gorm.DB) {

	userValidation := user.NewAdaptor(db)
	userCore := user.NewCore(db)
	userUsecase := user.NewUsecase(userCore, userValidation)
	userHandler := user.NewHandler(userUsecase)

	router.POST("/user", userHandler.NewUser)
	router.GET("/user/:id", userHandler.GetUser)
}

func UploadRouter(router *gin.Engine, client *storage.Client) {

	uploadCore := upload.NewCore(client)
	uploadUsecase := upload.NewUsecase(uploadCore)
	uploadHandler := upload.NewHandler(uploadUsecase)

	router.POST("/upload", uploadHandler.Uploadfile)
}


