package agentdetail

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewAgentDetail(c context.Context, agentDetail AgentDetail) error
	GetAgentDetails(c context.Context, firebaseId string) (*[]AgentDetailEntity, error)
	ListAgentDetails(c context.Context) (*[]AgentDetailEntity, error)
	GetAgentByID(c context.Context, id int) (*AgentDetailEntity, error)
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
		c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}
	agentDetail := AgentDetail{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Prompt:      req.Prompt,
		FirebaseID:      req.FirebaseID,
		FrameworkID: req.FrameworkID,
		RoleFrameID: req.RoleFrameID,
		TotalUsed: req.TotalUsed,
	}

	if err := h.usecase.NewAgentDetail(context.Background(), agentDetail); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"agents": agentDetails,
	})
}


func (h *Handler) GetAgentByID(c *gin.Context) {
    id := c.Param("id")
    roleID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid role ID",
        })
        return
    }
	// role_id := uint(roleID) 

    agent, err := h.usecase.GetAgentByID(context.Background(), roleID)
    if err != nil {
        c.AbortWithStatus(500)
        return
    }

    c.JSON(http.StatusOK, gin.H{"agent": agent})
}


func (h *Handler) ListAgentDetails(c *gin.Context) {
	agents, err := h.usecase.ListAgentDetails(context.Background())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"agents": agents})
}