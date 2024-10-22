package generate

import (
	"context"
	"github.com/gin-gonic/gin"
)
type Service interface {
	Generate(ctx context.Context, generateRequest GenerateRequest, language string) (string, error)
}

type GenerateHandler struct {
	service Service
}

func NewHandler(s Service) *GenerateHandler {
	return &GenerateHandler{service: s}
}

func (h *GenerateHandler) Generate(c *gin.Context) {
	var request GenerateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	language := c.Param("language")
	if language == "" {
        c.JSON(400, map[string]string{
            "error": "Language not set",
        })
        return
    }

	result, err := h.service.Generate(c.Request.Context(), request,language)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result == "" {
		c.JSON(500, gin.H{"error": "Generation failed"})
		return
	}

	c.JSON(201, gin.H{"result": result})
}
