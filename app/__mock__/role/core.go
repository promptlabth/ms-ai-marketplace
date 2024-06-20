package role

import (
	"context"

	"gorm.io/gorm"
)

// Core is responsible for the data storage operations related to roles.
type Core struct {
	db *gorm.DB
}

// NewCore creates a new instance of Core with a database connection.
func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

// CreateRole inserts a new role into the database.
func (c *Core) CreateRole(ctx context.Context, role RoleEntity) (*uint, error) {
	if err := c.db.Create(&role); err.Error != nil {
		return nil, err.Error
	}
	return &role.ID, nil
}

// GetRoleByID retrieves a role by their ID from the database.
func (c *Core) GetRoleByID(ctx context.Context, id uint) (*RoleEntity, error) {
	var role RoleEntity
	if err := c.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

//GET List of Roles
func (c *Core) ListRoles(ctx context.Context,language string) (*[]RoleEntity, error) {
	var roles []RoleEntity
	query := c.db.WithContext(ctx).Where("language = ?", language).Find(&roles)
	if query.Error != nil {
		return nil, query.Error
	}
	return &roles, nil
}


// UpdateRole updates a role's information in the database.
func (c *Core) UpdateRole(ctx context.Context, role RoleEntity) error {
	return c.db.Model(RoleEntity{}).Where("id = ?", role.ID).Updates(role).Error
}

// DeleteRole removes a role from the database by their ID.
func (c *Core) DeleteRole(ctx context.Context, id uint) error {
	return c.db.Delete(&RoleEntity{}, id).Error
}
