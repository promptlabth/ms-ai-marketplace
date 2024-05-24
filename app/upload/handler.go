package upload

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	// Check the content type
	if c.ContentType() != "multipart/form-data" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Content-Type must be multipart/form-data",
		})
		return
	}

	// Retrieve the file from the form
	f, err := c.FormFile("file_input")
	if err != nil {
		logrus.Errorf("Failed to retrieve file from form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve file from form",
		})
		return
	}

	// Open the file
	blobFile, err := f.Open()
	if err != nil {
		logrus.Errorf("Failed to open file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer blobFile.Close()

	// Upload the file using the usecase
	ctx := c.Request.Context()
	url, err := h.usecase.Uploadfile(ctx, blobFile, f.Filename)
	if err != nil {
		logrus.Errorf("Failed to upload file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload file",
		})
		return
	}

	// Respond with the URL of the uploaded file
	c.JSON(http.StatusCreated, gin.H{
		"success": "success",
		"url":     *url,
	})
}
