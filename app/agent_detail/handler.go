package agentdetail

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewAgentDetail(c context.Context, agentDetail AgentDetail) error
	GetAgentDetails(c context.Context, firebaseId string) (*[]AgentDetailEntity, error)
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) NewAgentDetail(c *gin.Context) {
	var req NewAgentDetailRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(404, map[string]string{
			"error": err.Error(),
		})
		return
	}
	agentDetail := AgentDetail{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Prompt:      req.Prompt,
		UserID:      req.UserID,
		FrameworkID: req.FrameworkID,
		RoleFrameID: req.RoleFrameID,
	}

	if err := h.usecase.NewAgentDetail(context.Background(), agentDetail); err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "creation sussess",
	})
}

func (h *Handler) GetAgentDetails(c *gin.Context) {
	firebaseID := c.Param("id")

	agentDetails, err := h.usecase.GetAgentDetails(c.Request.Context(), firebaseID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"agents": agentDetails,
	})
}
