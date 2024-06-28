package styleprompt

import (
	"context"

	"gorm.io/gorm"
)

// Core is responsible for the data storage operations related to style prompts.
type Core struct {
	db *gorm.DB
}

// NewCore creates a new instance of Core with a database connection.
func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

// CreateStylePrompt inserts a new style prompt into the database.
func (c *Core) CreateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) (uint, error) {
	if err := c.db.Create(&stylePrompt).Error; err != nil {
		return 0, err
	}
	return stylePrompt.ID, nil
}

// GetStylePromptByID retrieves a style prompt by its ID from the database.
func (c *Core) GetStylePromptByID(ctx context.Context, id uint,language string) (*StylePromptEntity, error) {
	var stylePrompt StylePromptEntity
	if err := c.db.WithContext(ctx).Where("language = ?", language).First(&stylePrompt, id).Error; err != nil {
		return nil, err
	}
	return &stylePrompt, nil
}

// ListStylePrompts gets a list of style prompts filtered by language from the database.
func (c *Core) ListStylePrompts(ctx context.Context, language string) (*[]StylePromptEntity, error) {
	var stylePrompts []StylePromptEntity
	query := c.db.WithContext(ctx).Where("language = ?", language).Find(&stylePrompts)
	if query.Error != nil {
		return nil, query.Error
	}
	return &stylePrompts, nil
}

// // UpdateStylePrompt updates a style prompt's information in the database.
// func (c *Core) UpdateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) error {
// 	return c.db.Model(&StylePromptEntity{}).Where("id = ?", stylePrompt.ID).Updates(stylePrompt).Error
// }

// // DeleteStylePrompt removes a style prompt from the database by its ID.
// func (c *Core) DeleteStylePrompt(ctx context.Context, id uint) error {
// 	return c.db.Delete(&StylePromptEntity{}, id).Error
// }
