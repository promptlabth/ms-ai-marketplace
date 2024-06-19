package framework

import (
	"context"

	"gorm.io/gorm"
)

// Core is responsible for the data storage operations related to frameworks.
type Core struct {
	db *gorm.DB
}

// NewCore creates a new instance of Core with a database connection.
func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

// CreateFramework inserts a new framework into the database.
func (c *Core) CreateFramework(ctx context.Context, framework FrameworkEntity) (*string, error) {
	if err := c.db.Create(&framework).Error; err != nil {
		return nil, err
	}
	return &framework.Name, nil
}

// GetFrameworkByID retrieves a framework by their ID from the database.
func (c *Core) GetFrameworkByID(ctx context.Context, id int) (*FrameworkEntity, error) {
	var framework FrameworkEntity
	if err := c.db.First(&framework, id).Error; err != nil {
		return nil, err
	}
	return &framework, nil
}

// ListFrameworks retrieves  frameworks by their ID from the database.
func (c *Core) ListFrameworks(ctx context.Context, language string) (*[]FrameworkEntity, error) {
	var frameworks []FrameworkEntity
	query := c.db.WithContext(ctx).Where("language = ?", language).Find(&frameworks)
	if query.Error != nil {
		return nil, query.Error
	}
	return &frameworks, nil
}

// UpdateFramework updates a framework's information in the database.
func (c *Core) UpdateFramework(ctx context.Context, framework FrameworkEntity) error {
	return c.db.Model(&FrameworkEntity{}).Where("id = ?", framework.ID).Updates(framework).Error
}

// DeleteFramework removes a framework from the database by their ID.
func (c *Core) DeleteFramework(ctx context.Context, id string) error {
	return c.db.Delete(&FrameworkEntity{}, "id = ?", id).Error
}
