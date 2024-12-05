package review

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

func (c *Core) CreateReview(ctx context.Context, review ReviewEntity) (*int, error) {

	if err := c.db.Create(&review); err.Error != nil {
		return nil, err.Error
	}
	return &review.ID, nil
}