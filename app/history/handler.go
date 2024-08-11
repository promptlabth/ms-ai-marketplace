package history

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
)
type usecase interface {
	CreateHistory(ctx context.Context, history History) (error)
}
type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}
func (h *Handler) GenerateMessage(c *gin.Context) {

	var req NewHistoryRequest

	ctx := c.Request.Context()

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	language := c.GetString("language")
	if language == "" {
		c.JSON(400, map[string]string{
			"error": "Language not set",
		})
		return
	}

	history := History{
		FirebaseID:     req.FirebaseID,
		AgentID:        req.AgentID,
		FrameworkID:    req.FrameworkID,
		Prompt:         req.Prompt,
		StyleMessageID: req.StyleMessageID,
		Language:       language,
	}

	if err := h.usecase.CreateHistory(ctx, history);  err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "CreateHistory successfully"})
}
