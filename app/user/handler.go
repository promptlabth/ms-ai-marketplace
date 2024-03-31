package user

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-orch-user-service/app"
)

type usecase interface {
	NewUser(c context.Context, user User) error
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) NewUser(c app.Context) {
	var req NewUserRequest

	ginCtx, ok := c.(app.Context)
	if !ok {
		c.AbortWithStatus(204)
		return
	}

	if err := ginCtx.Bind(&req); err != nil {
		c.BadRequest(err)
		return
	}

	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.usecase.NewUser(context.Background(), user); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.OK(gin.H{"message": "User created successfully"})


}
