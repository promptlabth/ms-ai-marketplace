package main

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	"github.com/promptlabth/ms-ai-marketplace/app/framework"
	"github.com/promptlabth/ms-ai-marketplace/app/generate"
	"github.com/promptlabth/ms-ai-marketplace/app/history"
	"github.com/promptlabth/ms-ai-marketplace/app/review"
	"github.com/promptlabth/ms-ai-marketplace/app/realtime_gen"
	styleprompt "github.com/promptlabth/ms-ai-marketplace/app/style_prompt"
	"github.com/promptlabth/ms-ai-marketplace/app/user/handler"
	"github.com/promptlabth/ms-ai-marketplace/app/user/repository"
	"github.com/promptlabth/ms-ai-marketplace/app/user/service"
	"github.com/promptlabth/ms-ai-marketplace/auth"
	"github.com/promptlabth/ms-ai-marketplace/config"
	"github.com/promptlabth/ms-ai-marketplace/middleware"
	"github.com/promptlabth/ms-ai-marketplace/app/payment/coins"
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
	// Initialize coins usecase
	coinsAdaptor := coins.NewAdaptor(db)
	coinsCore := coins.NewCore(db)
	coinsUsecase := coins.NewUsecase(coinsCore, coinsAdaptor)

	agentDetailValidation := agentdetail.NewAdaptor(db)
	agentDetailCore := agentdetail.NewCore(db)
	agentDetailUsecase := agentdetail.NewUsecase(agentDetailCore, agentDetailValidation, *coinsUsecase)
	agentDetailHandler := agentdetail.NewHandler(agentDetailUsecase)

	protected := router.Group("/creator")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/agent_detail", agentDetailHandler.NewAgentDetail)
	protected.PATCH("/update_agent/:id", agentDetailHandler.UpdateAgentDetail)
	protected.GET("/agent/user_id", agentDetailHandler.GetAgentDetails)
	protected.GET("/agents", agentDetailHandler.ListAgentDetails)
	router.GET("/creator/agents/approve", agentDetailHandler.ListAgentDetailsThatApprove)
	protected.GET("/agent/:id", agentDetailHandler.GetAgentByID)

	router.GET("/customer/:id", agentDetailHandler.GetAgentByID)
	router.POST("/customer/increase_used/:agent_id", agentDetailHandler.IncrementTotalUsed)
	router.POST("/admin/:agent_id/:status", agentDetailHandler.UpdateAgentStatus)
}

func FrameworkRouter(router *gin.Engine, db *gorm.DB) {
	frameworkValidation := framework.NewAdaptor(db)
	frameworkCore := framework.NewCore(db)
	frameworkUsecase := framework.NewUsecase(frameworkCore, frameworkValidation)
	frameworkHandler := framework.NewHandler(frameworkUsecase)

	protected := router.Group("/creator")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/framework", frameworkHandler.NewFramework)
	protected.GET("/frameworks/:language", frameworkHandler.ListFrameworks)
	protected.GET("/framework/:id", frameworkHandler.GetFrameworkByID)
}

func UsersRouter(router *gin.Engine, db *gorm.DB) {
	userRepositoryDB := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepositoryDB)
	userHandler := handler.NewUserHandler(userService)

	router.POST("/users/login", userHandler.LoginHandler)

	protected := router.Group("/users")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/me", userHandler.GetUserByFirebaseID)
}

func RoleRouter(router *gin.Engine, db *gorm.DB) {
	roleValidation := role.NewAdaptor(db)
	roleCore := role.NewCore(db)
	roleUsecase := role.NewUsecase(roleCore, roleValidation)
	roleHandler := role.NewHandler(roleUsecase)

	protected := router.Group("/creator")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/role", roleHandler.NewRole)
	protected.GET("/roles/:language", roleHandler.ListRoles)
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

	protected := router.Group("/creator")
	protected.Use(middleware.JWTMiddleware())
	router.POST("/creator/upload", uploadHandler.Uploadfile)
}

func StylePromptRouter(router *gin.Engine, db *gorm.DB) {
	stylePromptValidation := styleprompt.NewAdaptor(db)
	stylePromptCore := styleprompt.NewCore(db)
	stylePromptUsecase := styleprompt.NewUsecase(stylePromptCore, stylePromptValidation)
	stylePromptHandler := styleprompt.NewHandler(stylePromptUsecase)

	router.GET("/customer/style_prompts/:language", stylePromptHandler.ListStylePrompts)
	router.GET("/customer/style_prompt/:id", stylePromptHandler.GetStylePromptByID)
}

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
	protected := router.Group("/customer")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/use_agent/messages/:language", generateHandler.Generate)
}

func CustomerGetListsAgentUsage(router *gin.Engine, db *gorm.DB) {
	agentHistoryUsageValidation := history.NewAdaptor(db)
	agentHistoryUsageCore := history.NewCore(db)
	agentHistoryUsageUsecase := history.NewUsecase(agentHistoryUsageCore, agentHistoryUsageValidation)
	agentHistoryUsageHandler := history.NewHandler(agentHistoryUsageUsecase)

	protected := router.Group("/customer")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/agent_usage", agentHistoryUsageHandler.GetHistoryByFirebaseID)
}

func RealtimeGenCreateHistory(router *gin.Engine, db *gorm.DB) {
	agentHistoryUsageValidation := history.NewAdaptor(db)
	agentHistoryUsageCore := history.NewCore(db)
	agentHistoryUsageUsecase := history.NewUsecase(agentHistoryUsageCore, agentHistoryUsageValidation)
	agentHistoryUsageHandler := history.NewHandler(agentHistoryUsageUsecase)

	protected := router.Group("/customer")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/create_history/:language", agentHistoryUsageHandler.CreateHistoryByFirebaseID)
}

func RealtimeGenGetFullPromptByAgentID(router *gin.Engine, db *gorm.DB) {
	realtimeGenCore := realtimegen.NewCore(db)
	realtimeGenUsecase := realtimegen.NewUsecase(realtimeGenCore)
	realtimeGenHandler := realtimegen.NewHandler(realtimeGenUsecase)

	router.GET("/customer/get_full_prompt/:agent_id", realtimeGenHandler.GetFullPromptByAgentID)
}

func ReviewRouter(router *gin.Engine, db *gorm.DB) {
	reviewCore := review.NewCore(db)
	agentDetailCore := agentdetail.NewCore(db)
	reviewUsecase := review.NewUsecase(reviewCore, agentDetailCore)
	reviewHandler := review.NewHandler(reviewUsecase)

	protected := router.Group("/review")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/", reviewHandler.NewReview)
	protected.GET("/latest/:agent_id", reviewHandler.GetLatestReviewByAgentID)
}

func CoinsRouter(router *gin.Engine, db *gorm.DB) {
	coinsAdaptor := coins.NewAdaptor(db)
	coinsCore := coins.NewCore(db)
	coinsUsecase := coins.NewUsecase(coinsCore, coinsAdaptor)
	coinsHandler := coins.NewHandler(coinsUsecase)

	protected := router.Group("/coins")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/", coinsHandler.CreateCoins)
	protected.POST("/add", coinsHandler.AddCoins)
	protected.GET("/:agentID", coinsHandler.GetCoins)
	protected.POST("/reset/:agentID", coinsHandler.SetCoinsToZero)
}