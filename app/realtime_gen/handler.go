package realtimegen

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	GetFullPromptByAgentID(ctx context.Context, id int) (*RealTimeGenPrompt, error)
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) GetFullPromptByAgentID(c *gin.Context) {
	agentID, err := strconv.Atoi(c.Param("agent_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	prompt, err := h.usecase.GetFullPromptByAgentID(c.Request.Context(), agentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, prompt)
}