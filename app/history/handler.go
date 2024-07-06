package history

import (
	"context"
	"net/http"
	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Define the usecase interface for history operations
type usecase interface {
	CreateHistory(ctx context.Context, history History) (*string, error)
// 	GetHistoryByID(ctx context.Context, id int) (*History, error)
// 	ListHistories(ctx context.Context, userID int) (*[]History, error)
// 	UpdateHistory(ctx context.Context, history History) error
// 	DeleteHistory(ctx context.Context, id int) error
}

// Handler handles HTTP requests for history operations
type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

// CreateHistory creates a new history 
func (h *Handler) GenerateMessage(c *gin.Context) {
	var req NewHistoryRequest

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
		UserID:         req.UserID,
		AgentID:        req.AgentID,
		FrameworkID:    req.FrameworkID,
		Prompt:         req.Prompt,
		StyleMessageID: req.StyleMessageID,
		Language:       language,
		TimeStamp:      time.Now(),
	}

	result , err := h.usecase.CreateHistory(context.Background(), history); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": result})
}

// // GetHistoryByID retrieves a history record by ID
// func (h *Handler) GetHistoryByID(c *gin.Context) {
// 	id := c.Param("id")
// 	historyID, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, map[string]string{
// 			"error": "Invalid history ID",
// 		})
// 		return
// 	}

// 	history, err := h.usecase.GetHistoryByID(context.Background(), historyID)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"history": history})
// }

// // ListHistories lists history records by user ID
// func (h *Handler) ListHistories(c *gin.Context) {
// 	userIDStr := c.Query("user_id")
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, map[string]string{
// 			"error": "Invalid user ID",
// 		})
// 		return
// 	}

// 	histories, err := h.usecase.ListHistories(context.Background(), userID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to list histories",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":    "success",
// 		"histories": histories,
// 	})
// }

// // UpdateHistory updates an existing history record
// func (h *Handler) UpdateHistory(c *gin.Context) {
// 	var req History
// 	if err := c.Bind(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, map[string]string{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	if err := h.usecase.UpdateHistory(context.Background(), req); err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "History updated successfully"})
// }

// // DeleteHistory deletes a history record by ID
// func (h *Handler) DeleteHistory(c *gin.Context) {
// 	id := c.Param("id")
// 	historyID, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, map[string]string{
// 			"error": "Invalid history ID",
// 		})
// 		return
// 	}

// 	if err := h.usecase.DeleteHistory(context.Background(), historyID); err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "History deleted successfully"})
// }
