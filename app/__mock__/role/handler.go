package role

import (
	"context"
	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewRole(ctx context.Context, role Role) error
	ListRoles(ctx context.Context) (*[]RoleEntity, error)
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
		c.JSON(404, map[string]string{
			"error": err.Error(),
		})
		return
	}

	role := Role{
		ID:   req.ID,
		Name: req.Name,
	}

	if err := h.usecase.NewRole(context.Background(), role); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, gin.H{"message": "Role created successfully"})
}

// NewRequest for get List of Roles
func (h *Handler) ListRoles(c *gin.Context) {
	roles, err := h.usecase.ListRoles(context.Background())
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, gin.H{"roles": roles})
}
