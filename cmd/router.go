package main

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/framework"
	"github.com/promptlabth/ms-ai-marketplace/app/history"
	styleprompt "github.com/promptlabth/ms-ai-marketplace/app/style_prompt"
	"github.com/promptlabth/ms-ai-marketplace/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/promptlabth/ms-ai-marketplace/app/role"
	"github.com/promptlabth/ms-ai-marketplace/app/upload"
	"github.com/promptlabth/ms-ai-marketplace/app/user"
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

func UserRouter(router *gin.Engine, db *gorm.DB) error {

	creds := insecure.NewCredentials()
	cc, err := grpc.NewClient(config.Val.Adaptor.User.Url, grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}
	userClient := user.NewUserServiceClient(cc)
	userValidation := user.NewAdaptor(db)
	userCore := user.NewCore(db)

	userAdaptor := user.NewUserAdaptor(userClient)
	userUsecase := user.NewUsecase(userCore, userValidation, userAdaptor)
	userHandler := user.NewHandler(userUsecase)

	router.POST("/user", userHandler.NewUser)
	router.GET("/user/:id", userHandler.GetUser)
	return nil
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
	stylePromptUsecase := styleprompt.NewUsecase(stylePromptCore, stylePromptValidation)
	stylePromptHandler := styleprompt.NewHandler(stylePromptUsecase)

	router.GET("/:lang/customer/style_prompts", stylePromptHandler.ListStylePrompts)
	router.GET("/:lang/customer/style_prompt/:id", stylePromptHandler.GetStylePromptByID)
}
func GenerateMessageRouter(router *gin.Engine, db *gorm.DB) {
	generateMessageValidation := history.NewAdaptor(db)
	generateMessageCore := history.NewCore(db)
	agentdetailCore := agentdetail.NewCore(db)
	generateMessageUsecase := history.NewUsecase(generateMessageCore, generateMessageValidation, agentdetailCore)
	generateMessageHandler := history.NewHandler(generateMessageUsecase)

	router.POST("/:lang/customer/use_agent/messages", generateMessageHandler.GenerateMessage)
	// router.GET("/:lang/customer/style_prompt/:id", stylePromptHandler.GetStylePromptByID)
}
