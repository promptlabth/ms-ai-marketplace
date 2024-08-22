package user

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-ai-marketplace/app"
)

type userUsecaser interface {
	GetUser(c context.Context, firebase_id string) (*UserEntity, error)
	LoginService(ctx context.Context, req LoginRequestDomain) (*LoginResponseDomain, error)
}

type Handler struct {
	userUsecase userUsecaser
}

func NewHandler(u userUsecaser) *Handler {
	return &Handler{userUsecase: u}
}

func (h *Handler) GetUser(c *gin.Context) {
	firebaseID := c.Param("id")
	userByID, err := h.userUsecase.GetUser(c.Request.Context(), firebaseID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, app.Response[any]{
			Code:  5000,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, app.Response[UserEntity]{
		Code: 1000,
		Data: userByID,
	})
}
