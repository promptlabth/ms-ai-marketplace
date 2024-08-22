package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-ai-marketplace/app"
	"github.com/promptlabth/ms-ai-marketplace/logger"
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
	ctx := c.Request.Context()
	firebaseID := c.Param("id")
	userByID, err := h.userUsecase.GetUser(ctx, firebaseID)
	if err != nil {
		logger.Error(ctx, err.Error())
		c.JSON(http.StatusOK, app.Response[any]{
			Code:  5000,
			Error: err.Error(),
		})
		return
	}
	logger.Info(ctx, "Test")
	c.JSON(http.StatusOK, app.Response[UserEntity]{
		Code: 1000,
		Data: userByID,
	})
}
