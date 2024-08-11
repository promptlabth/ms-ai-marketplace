package role

import (
	"context"

	"gorm.io/gorm"
)

// Core is responsible for the data storage operations related to roles.
type Core struct {
	db *gorm.DB
}

func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

func (c *Core) CreateRole(ctx context.Context, role RoleEntity) (*int, error) {
	if err := c.db.Create(&role); err.Error != nil {
		return nil, err.Error
	}
	return &role.ID, nil
}

func (c *Core) GetRoleByID(ctx context.Context, id int) (*RoleEntity, error) {
	var role RoleEntity
	if err := c.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (c *Core) ListRoles(ctx context.Context,language string) (*[]RoleEntity, error) {
	var roles []RoleEntity
	query := c.db.WithContext(ctx).Where("language = ?", language).Find(&roles)
	if query.Error != nil {
		return nil, query.Error
	}
	return &roles, nil
}

func (c *Core) UpdateRole(ctx context.Context, role RoleEntity) error {
	return c.db.Model(RoleEntity{}).Where("id = ?", role.ID).Updates(role).Error
}

func (c *Core) DeleteRole(ctx context.Context, id int) error {
	return c.db.Delete(&RoleEntity{}, id).Error
}
