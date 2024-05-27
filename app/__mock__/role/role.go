package role

import (
	"context"
)

// Role represents the structure of a role in the system.
type Role struct {
	ID   int
	Name string
}

// RoleInterface defines the set of methods that any implementation of the Role service must provide.
type RoleInterface interface {
	CreateRole(ctx context.Context, role Role) (int64, error) // Creates a new role and returns the role ID
	GetRoleByID(ctx context.Context, id int64) (*Role, error) // Fetches a role by their ID
	UpdateRole(ctx context.Context, role Role) error          // Updates an existing role
	DeleteRole(ctx context.Context, id int64) error           // Deletes a role by their ID
	ListRoles(ctx context.Context) (*[]RoleEntity, error)
}

type NewRoleRequest struct {
	Name string `json:"name"`
}
