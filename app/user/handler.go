package user

import (
	"context"

	"github.com/gin-gonic/gin"
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

func (h *Handler) NewUser(c *gin.Context) {
	var req NewUserRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(404, map[string]string{
			"error": err.Error(),
		})
		return
	}

	user := User{
		FriebaseID: req.Username,
		Email:      req.Email,
		Password:   req.Password,
	}

	if err := h.usecase.NewUser(context.Background(), user); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})

}
