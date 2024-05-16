package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-orch-user-service/app/upload"
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


func UploadRouter(router *gin.Engine){
	t := "upload"
	uploadCore := upload.NewCore(&t)
	// if uploadCore == nil {
	// 	fmt.Println("Checking Uploadcore")
	// 	fmt.Printf("uploadCore: %+v\n", uploadCore)
	// 	return 
	// } else {
	// 	fmt.Println("uploadCore is nil")
	// }
	
	uploadUsecase := upload.NewUsecase(uploadCore)
	uploadHandler := upload.NewHandler(uploadUsecase)

	router.POST("/upload", uploadHandler.Uploadfile)
}