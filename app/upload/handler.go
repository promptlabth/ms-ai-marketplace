package upload

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	Uploadfile(ctx context.Context, file multipart.File, filename string) (*string, error)
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) Uploadfile(c *gin.Context) {
	f, err := c.FormFile("file_input")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	blobFile, err := f.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer blobFile.Close()

	ctx := c.Request.Context()
	url, err := h.usecase.Uploadfile(ctx, blobFile, f.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": "success",
		"url":     *url,
	})
}
