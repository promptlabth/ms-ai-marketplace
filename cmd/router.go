package main

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-orch-user-service/app/agent_detail"
	"github.com/promptlabth/ms-orch-user-service/app/framework"
	"github.com/promptlabth/ms-orch-user-service/app/history"
	styleprompt "github.com/promptlabth/ms-orch-user-service/app/style_prompt"

	// "github.com/promptlabth/ms-orch-user-service/app/role"
	"github.com/promptlabth/ms-orch-user-service/app/__mock__/role"
	"github.com/promptlabth/ms-orch-user-service/app/upload"
	"github.com/promptlabth/ms-orch-user-service/app/user"
	"gorm.io/gorm"
)

func AgentDetailRouter(router *gin.Engine, db *gorm.DB) {
	agentDetailValidation := agentdetail.NewAdaptor(db)
	agentDetailCore := agentdetail.NewCore(db)
	agentDetailUsecase := agentdetail.NewUsecase(agentDetailCore, agentDetailValidation)
	agentDetailHandler := agentdetail.NewHandler(agentDetailUsecase)

	router.POST("/creator/agent_detail", agentDetailHandler.NewAgentDetail)
	router.GET("/creator/agent/user_id/:id", agentDetailHandler.GetAgentDetails)
	router.GET("/creator/agents", agentDetailHandler.ListAgentDetails)
	router.GET("/creator/agent/:id", agentDetailHandler.GetAgentByID)
	router.GET("/customer/:id", agentDetailHandler.GetAgentByID)
}

func FrameworkRouter(router *gin.Engine, db *gorm.DB) {
	frameworkValidation := framework.NewAdaptor(db)
	frameworkCore := framework.NewCore(db)
	frameworkUsecase := framework.NewUsecase(frameworkCore, frameworkValidation)
	frameworkHandler := framework.NewHandler(frameworkUsecase)

	router.POST("/creator/framework", frameworkHandler.NewFramework)
	router.GET("/:lang/creator/frameworks", frameworkHandler.ListFrameworks)
	router.GET("creator/framework/:id", frameworkHandler.GetFrameworkByID)
}

func RoleRouter(router *gin.Engine, db *gorm.DB) {

	roleValidation := role.NewAdaptor(db)
	roleCore := role.NewCore(db)
	roleUsecase := role.NewUsecase(roleCore, roleValidation)
	roleHandler := role.NewHandler(roleUsecase)

	router.POST("/creator/role", roleHandler.NewRole)
	router.GET("/:lang/creator/roles", roleHandler.ListRoles)
	router.GET("/creator/role/:id", roleHandler.GetRoleByID)
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

	router.POST("/creator/upload", uploadHandler.Uploadfile)
}

func StylePromptRouter(router *gin.Engine, db *gorm.DB) {
	stylePromptValidation := styleprompt.NewAdaptor(db)
	stylePromptCore := styleprompt.NewCore(db)
	stylePromptUsecase := styleprompt.NewUsecase(stylePromptCore,stylePromptValidation)
	stylePromptHandler := styleprompt.NewHandler(stylePromptUsecase)

	router.GET("/:lang/customer/style_prompts", stylePromptHandler.ListStylePrompts)
	router.GET("/:lang/customer/style_prompt/:id", stylePromptHandler.GetStylePromptByID)
}
func GenerateMessageRouter(router *gin.Engine, db *gorm.DB) {
	generateMessageValidation := history.NewAdaptor(db)
	generateMessageCore := history.NewCore(db)
	generateMessageUsecase := history.NewUsecase(generateMessageCore,generateMessageValidation)
	generateMessageHandler := history.NewHandler(generateMessageUsecase)

	router.POST("/:lang/customer/use_agent/messages", generateMessageHandler.GenerateMessage)
	// router.GET("/:lang/customer/style_prompt/:id", stylePromptHandler.GetStylePromptByID)
}

