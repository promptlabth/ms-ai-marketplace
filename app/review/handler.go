package review

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewReview(ctx context.Context, review ReviewEntity) error
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) NewReview(c *gin.Context) {
	var req ReviewRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}

	review := ReviewEntity{
		AdminID:  req.AdminID,
		AgentID:  req.AgentID,
		Reason:   req.Reason,
		DateTime: req.DateTime,
	}

	if err := h.usecase.NewReview(context.Background(), review); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "creation success",
	})
}
