package framework

import (
	"context"

	"gorm.io/gorm"
)
type Core struct {
	db *gorm.DB
}

func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

func (c *Core) CreateFramework(ctx context.Context, framework FrameworkEntity) (*string, error) {
	if err := c.db.Create(&framework).Error; err != nil {
		return nil, err
	}
	return &framework.Name, nil
}

func (c *Core) GetFrameworkByID(ctx context.Context, id int) (*FrameworkEntity, error) {
	var framework FrameworkEntity
	if err := c.db.First(&framework, id).Error; err != nil {
		return nil, err
	}
	return &framework, nil
}

func (c *Core) ListFrameworks(ctx context.Context, language string) (*[]FrameworkEntity, error) {
	var frameworks []FrameworkEntity
	query := c.db.WithContext(ctx).Where("language = ?", language).Find(&frameworks)
	if query.Error != nil {
		return nil, query.Error
	}
	return &frameworks, nil
}

func (c *Core) UpdateFramework(ctx context.Context, framework FrameworkEntity) error {
	return c.db.Model(&FrameworkEntity{}).Where("id = ?", framework.ID).Updates(framework).Error
}

func (c *Core) DeleteFramework(ctx context.Context, id string) error {
	return c.db.Delete(&FrameworkEntity{}, "id = ?", id).Error
}
