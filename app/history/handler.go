package history

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	// "golang.org/x/text/internal/language"
)

type usecase interface {
	CreateHistory(ctx context.Context, history History) error
	GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error)
	GetHistoriesByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error)
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
		FirebaseID:        req.FirebaseID,
		AgentID:           req.AgentID,
		FrameworkID:       req.FrameworkID,
		Prompt:            req.Prompt,
		StyleMessageID:    req.StyleMessageID,
		Language:          language,
		Result:            req.Result,
		Model:             req.Model,
	}

	if err := h.usecase.CreateHistory(ctx, history); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "CreateHistory successfully"})
}

func (h *Handler) GetHistoryByFirebaseID(c *gin.Context) {
	firebaseID, exists := c.Get("firebase_id")
	if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Firebase ID not found in token"})
        return
    }

    // Type assertion to convert firebaseID from any to string
    firebaseIDStr, ok := firebaseID.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Firebase ID type"})
        return
    }

	histories, err := h.usecase.GetHistoryByFirebaseID(c.Request.Context(), firebaseIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, histories)
}

func (h *Handler) GetHistoriesByFirebaseID(c *gin.Context) {
    firebaseID, exists := c.Get("firebase_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Firebase ID not found in token"})
        return
    }

    // Type assertion to convert firebaseID from any to string
    firebaseIDStr, ok := firebaseID.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Firebase ID type"})
        return
    }

    histories, err := h.usecase.GetHistoriesByFirebaseID(c.Request.Context(), firebaseIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, map[string]string{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, histories)
}

func (h *Handler) CreateHistoryByFirebaseID(c *gin.Context) {
	firebaseID, exists := c.Get("firebase_id")
	if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Firebase ID not found in token"})
        return
    }
	
	language := c.Param("language")
	var req NewHistoryRequest

	ctx := c.Request.Context()

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	history := History{
		FirebaseID:        firebaseID.(string),
		Language:          language,
		AgentID:           req.AgentID,
		FrameworkID:       req.FrameworkID,
		Prompt:            req.Prompt,
		StyleMessageID:    req.StyleMessageID,
		Result:            req.Result,
		Model:             req.Model,
	}
	if err := h.usecase.CreateHistory(ctx, history); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "CreateHistory successfully"})

}
