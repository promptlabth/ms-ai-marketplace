package main

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/framework"
	"github.com/promptlabth/ms-ai-marketplace/app/generate"
	"github.com/promptlabth/ms-ai-marketplace/app/history"
	styleprompt "github.com/promptlabth/ms-ai-marketplace/app/style_prompt"
	"github.com/promptlabth/ms-ai-marketplace/auth"
	"github.com/promptlabth/ms-ai-marketplace/config"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/mock/gomock"

	"github.com/promptlabth/ms-ai-marketplace/app/role"
	"github.com/promptlabth/ms-ai-marketplace/app/upload"
	"github.com/promptlabth/ms-ai-marketplace/app/user"
	userProto "github.com/promptlabth/proto-lib/user"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
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

func UserRouter(ctx context.Context, router *gin.Engine, db *gorm.DB) error {

	cc, err := InitialGRpc(config.Val.Adaptor.User.Url)
	if err != nil {
		return err
	}
	userClient := userProto.NewUserServiceClient(cc)

	userCore := user.NewCore(db)

	app, err := auth.Init()
	if err != nil {
		return err
	}

	grpcUserServer := user.NewGrpcRequestor(userClient)
	userAdaptor := user.NewUserAdaptor(app)
	userUsecase := user.NewUsecase(userCore, userAdaptor, grpcUserServer)
	userHandler := user.NewHandler(userUsecase)

	user := router.Group("/user")
	otelopts := otelgin.WithPropagators(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	user.Use(
		otelgin.Middleware("ms-ai-marketplace", otelopts),
		LoggingWithDumbBody(),
	)
	user.POST("/login", userHandler.LoginHandler)
	user.GET("/:id", userHandler.GetUser)
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
	router.GET("/customer/style_prompt/:id", stylePromptHandler.GetStylePromptByID)
}

// func GenerateMessageRouter(router *gin.Engine, db *gorm.DB) {
// 	generateMessageValidation := history.NewAdaptor(db)
// 	generateMessageCore := history.NewCore(db)
// 	agentdetailCore := agentdetail.NewCore(db)
// 	generateMessageUsecase := history.NewUsecase(generateMessageCore, generateMessageValidation, agentdetailCore)
// 	generateMessageHandler := history.NewHandler(generateMessageUsecase)

//		router.POST("/:lang/customer/use_agent/messages", generateMessageHandler.GenerateMessage)
//		// router.GET("/:lang/customer/style_prompt/:id", stylePromptHandler.GetStylePromptByID)
//	}
func GenerateMessageRouter(router *gin.Engine, db *gorm.DB, ctrl *gomock.Controller) {
	// Initialize mocks
	generateAdaptor := generate.NewMockgenerateAdaptor(ctrl)
	agentDetailCore := agentdetail.NewCore(db)
	stylePromptCore := styleprompt.NewCore(db)
	frameworkCore := framework.NewCore(db)
	roleCore := role.NewCore(db)
	historyCore := history.NewCore(db)
	generateCore := generate.NewCore(db)

	// Initialize the service with actual implementations and mock
	generateService := generate.NewService(
		generateAdaptor,
		agentDetailCore,
		stylePromptCore,
		frameworkCore,
		roleCore,
		historyCore,
		generateCore,
	)

	// Initialize the handler
	generateHandler := generate.NewHandler(generateService)

	// Define routes and handlers
	router.POST("/:lang/customer/use_agent/messages", generateHandler.Generate)
}
