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
	"github.com/promptlabth/ms-orch-user-service/config"
	"github.com/promptlabth/ms-orch-user-service/database"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	db := database.NewGormDBWithDefault()

	// initial storage bucket
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("prompt-lab-383408-512938be4baf.json"))
	if err != nil {
		log.Fatal(err)
	}
	// r := app.NewRouter(logger)
	// r := app.NewRouterGin(logger)
	r := gin.Default()
	// r.Use(gin.WrapF(cors.New(opts).HandlerFunc))

	r.Use(CORSMiddleware())
	UserRouter(r, db)
	AgentDetailRouter(r, db)
	UploadRouter(r, client)

	srv := http.Server{
		Addr:              ":" + config.Val.Port,
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
