package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-orch-user-service/config"
	"github.com/promptlabth/ms-orch-user-service/database"
	"github.com/promptlabth/ms-orch-user-service/logger"
)

func main() {

	db := database.NewGormDBWithDefault()

	logger := logger.New()
	// r := app.NewRouter(logger)
	// r := app.NewRouterGin(logger)
	r := gin.Default()
	// r.Use(gin.WrapF(cors.New(opts).HandlerFunc))

	r.Use(CORSMiddleware())
	NewRouter(r, db)

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
			logger.Info("HTTP server Shutdown: " + err.Error())
		}
		close(idleConnsClosed)
	}()

	fmt.Println(":" + config.Val.Port + " is serve")

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("HTTP server ListenAndServe: " + err.Error())
		return
	}

	<-idleConnsClosed
	fmt.Println("gracefully")

}
