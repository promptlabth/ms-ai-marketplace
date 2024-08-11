package styleprompt

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

func (c *Core) CreateStylePrompt(ctx context.Context, stylePrompt StylePromptEntity) (int, error) {
	if err := c.db.Create(&stylePrompt).Error; err != nil {
		return 0, err
	}
	return stylePrompt.ID, nil
}

func (c *Core) GetStylePromptByID(ctx context.Context, id int) (*StylePromptEntity, error) {
	var stylePrompt StylePromptEntity
	if err := c.db.First(&stylePrompt, id).Error; err != nil {
		return nil, err	
	}
	return &stylePrompt, nil
}

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
// func (c *Core) DeleteStylePrompt(ctx context.Context, id int) error {
// 	return c.db.Delete(&StylePromptEntity{}, id).Error
// }
