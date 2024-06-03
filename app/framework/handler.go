package framework

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewFramework(ctx context.Context, framework Framework) error
	ListFrameworks(ctx context.Context) (*[]FrameworkEntity, error)
	GetFrameworkByID(ctx context.Context, id int) (*FrameworkEntity, error)
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) NewFramework(c *gin.Context) {
	var req NewFrameworkRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(404, map[string]string{
			"error": err.Error(),
		})
		return
	}

	framework := Framework{
		Name:      req.Name,
		Detail:    req.Detail,
		Component: req.Component,
	}

	if err := h.usecase.NewFramework(context.Background(), framework); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Framework created successfully"})
}



func (h *Handler) GetFrameworkByID(c *gin.Context) {
	id := c.Param("id")
	framwork_id, err := strconv.Atoi(id)
	if err != nil {
        c.JSON(400, map[string]string{
            "error": "Invalid role ID",
        })
        return
    }
	// role_id := uint(roleID) 

    framework, err := h.usecase.GetFrameworkByID(context.Background(), framwork_id)
    if err != nil {
        c.AbortWithStatus(500)
        return
    }

    c.JSON(http.StatusOK, gin.H{"framework": framework})
}

func (h *Handler) ListFrameworks(c *gin.Context) {

	frameworks, err := h.usecase.ListFrameworks(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"frameworks": frameworks,
	})
}
