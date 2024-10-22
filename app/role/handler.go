package role

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewRole(ctx context.Context, role Role) error
	ListRoles(ctx context.Context,language string) (*[]RoleEntity, error)
	GetRoleByID(ctx context.Context, id int) (*RoleEntity, error)
	UpdateRole(ctx context.Context, role RoleEntity) error
	DeleteRole(ctx context.Context, id int) error
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) NewRole(c *gin.Context) {
	var req NewRoleRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}

	role := Role{
		Name: req.Name,
	}

	if err := h.usecase.NewRole(context.Background(), role); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role created successfully"})
}


func (h *Handler) ListRoles(c *gin.Context) {

	language := c.Param("language")
	if language == "" {
        c.JSON(400, map[string]string{
            "error": "Language not set",
        })
        return
    }
	roles, err := h.usecase.ListRoles(context.Background(),language)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

func (h *Handler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, map[string]string{
			"error": "Invalid role ID",
		})
		return
	}
	role_id := roleID

	role, err := h.usecase.GetRoleByID(context.Background(), role_id)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

func (h *Handler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, map[string]string{
			"error": "Invalid role ID",
		})
		return
	}
	role_id := roleID
	if err := h.usecase.DeleteRole(context.Background(), role_id); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
