package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-ai-marketplace/config"
	"github.com/promptlabth/ms-ai-marketplace/database"
	"github.com/promptlabth/ms-ai-marketplace/logger"
	"go.uber.org/mock/gomock"

	// "github.com/promptlabth/ms-ai-marketplace/initializers"
	"google.golang.org/api/option"
)

func main() {

	// load .env file if ENV == local
	// initializers.LoadEnvVariables()
	ctx := context.Background()
	logger.InitLogger()
	defer logger.Sync()

	// init trace for otel
	tp, err := config.InitTrace()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	db := database.NewGormDBWithDefault()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("prompt-lab-cred.json"))
	if err != nil {
		logger.Fatal(ctx, err.Error())
	}

	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()
	// r := app.NewRouter(logger)
	// r := app.NewRouterGin(logger)
	r := gin.Default()
	// r.Use(gin.WrapF(cors.New(opts).HandlerFunc))

	r.Use(CORSMiddleware())
	AgentDetailRouter(r, db)
	FrameworkRouter(r, db)
	RoleRouter(r, db)
	if err := UserRouter(ctx, r, db); err != nil {
		logger.Fatal(ctx, err.Error())
	}
	StylePromptRouter(r, db)
	UploadRouter(r, client)
	GenerateMessageRouter(r, db, ctrl)

	port := config.Val.Port
	if port == "" {
		fmt.Println("use a default port :8080")
		port = "8080" // Default port if not specified
	}
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	srv := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		d := time.Duration(5 * time.Second)
		fmt.Printf("shutting down int %s ...", d)
		// We received an interrupt signal, shut down.
		ctx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Fatal(err)
		}
		close(idleConnsClosed)
	}()

	fmt.Println(":" + config.Val.Port + " is serve")

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
		return
	}

	<-idleConnsClosed
	fmt.Println("gracefully")
}
