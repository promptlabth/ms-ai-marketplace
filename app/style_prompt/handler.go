package styleprompt

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type stylePromptUsecase interface {
	CreateStylePrompt(ctx context.Context, stylePrompt StylePrompt) (int, error)
	ListStylePrompts(ctx context.Context, language string) (*[]StylePromptEntity, error)
	GetStylePromptByID(ctx context.Context, id int) (*StylePromptEntity, error)
	// UpdateStylePrompt(ctx context.Context, stylePrompt StylePrompt) error
	// DeleteStylePrompt(ctx context.Context, id int) error
}

type Handler struct {
	usecase stylePromptUsecase
}

func NewHandler(u stylePromptUsecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) CreateStylePrompt(c *gin.Context) {
	var req NewStylePromptRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	stylePrompt := StylePrompt{
		Name:     req.Name,
		Language: req.Language,
	}

	id, err := h.usecase.CreateStylePrompt(context.Background(), stylePrompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create style prompt",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Style prompt created successfully",
		"id":      id,
	})
}


func (h *Handler) ListStylePrompts(c *gin.Context) {

	language := c.Param("language")
	if language == "" {
        c.JSON(400, map[string]string{
            "error": "Language not set",
        })
        return
    }
	data, err := h.usecase.ListStylePrompts(context.Background(),language)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) GetStylePromptByID(c *gin.Context) {
	id := c.Param("id")
	stylePromptID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, map[string]string{
			"error": "Invalid stylePrompt ID",
		})
		return
	}
	stylePrompt_id :=stylePromptID

	data, err := h.usecase.GetStylePromptByID(context.Background(), stylePrompt_id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
