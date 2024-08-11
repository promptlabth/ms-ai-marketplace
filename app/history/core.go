package history

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

func (c *Core) CreateHistory(ctx context.Context, history HistoryEntity) (*int, error) {
	if err := c.db.Create(&history).Error; err != nil {
		return nil, err
	}
	return &history.ID, nil
}
