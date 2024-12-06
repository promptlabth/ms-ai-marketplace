package review

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewReview(ctx context.Context, review ReviewEntity) error
	GetLatestReviewByAgentID(ctx context.Context, agentID int) (*ReviewEntity, error)
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
		DateTime: time.Now(),
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

func (h *Handler) GetLatestReviewByAgentID(c *gin.Context) {
	agentID, err := strconv.Atoi(c.Param("agent_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID"})
		return
	}

	review, err := h.usecase.GetLatestReviewByAgentID(context.Background(), agentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get latest review"})
		return
	}

	c.JSON(http.StatusOK, review)
}
