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

	// "github.com/promptlabth/ms-ai-marketplace/initializers"
	"google.golang.org/api/option"
)

func main() {

	// load .env file if ENV == local
	// initializers.LoadEnvVariables()
	ctx := context.Background()

	db := database.NewGormDBWithDefault()

	logx, stop := logger.NewZap()
	defer stop()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("prompt-lab-cred.json"))
	if err != nil {
		logx.Fatal(err.Error())
	}
	// r := app.NewRouter(logger)
	// r := app.NewRouterGin(logger)
	r := gin.Default()
	// r.Use(gin.WrapF(cors.New(opts).HandlerFunc))

	r.Use(CORSMiddleware())
	AgentDetailRouter(r, db)
	FrameworkRouter(r, db)
	RoleRouter(r, db)
	UserRouter(r, db)
	StylePromptRouter(r, db)
	GenerateMessageRouter(r, db)
	UploadRouter(r, client)

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
