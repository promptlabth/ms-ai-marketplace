package history

import (
	"context"
	"log"
	// "fmt"
	"net/http"

	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Define the usecase interface for history operations
type usecase interface {
	CreateHistory(ctx context.Context, history History) (*string, string)
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
		FirebaseID:         req.FirebaseID,
		AgentID:        req.AgentID,
		FrameworkID:    req.FrameworkID,
		Prompt:         req.Prompt,
		StyleMessageID: req.StyleMessageID,
		Language:       language,
		TimeStamp:      time.Now(),
	}

	result , err := h.usecase.CreateHistory(context.Background(), history); 
	if err != "" {
		log.Print("CreateHistory"+err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CreateHistory"+err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": result})
}
