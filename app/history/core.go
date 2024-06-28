package history

import (
	"context"

	"gorm.io/gorm"
)

// Core is responsible for the data storage operations related to histories.
type Core struct {
	db *gorm.DB
}

// NewCore creates a new instance of Core with a database connection.
func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

// CreateHistory inserts a new history record into the database.
func (c *Core) CreateHistory(ctx context.Context, history HistoryEntity) (*int, error) {
	if err := c.db.Create(&history).Error; err != nil {
		return nil, err
	}
	return &history.ID, nil
}

// // GetHistoryByID retrieves a history record by its ID from the database.
// func (c *Core) GetHistoryByID(ctx context.Context, id int) (*History, error) {
// 	var history History
// 	if err := c.db.First(&history, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &history, nil
// }

// // ListHistories retrieves history records by user ID from the database.
// func (c *Core) ListHistories(ctx context.Context, userID int) (*[]History, error) {
// 	var histories []History
// 	query := c.db.WithContext(ctx).Where("user_id = ?", userID).Find(&histories)
// 	if query.Error != nil {
// 		return nil, query.Error
// 	}
// 	return &histories, nil
// }

// // UpdateHistory updates a history record's information in the database.
// func (c *Core) UpdateHistory(ctx context.Context, history History) error {
// 	return c.db.Model(&History{}).Where("id = ?", history.ID).Updates(history).Error
// }

// // DeleteHistory removes a history record from the database by its ID.
// func (c *Core) DeleteHistory(ctx context.Context, id int) error {
// 	return c.db.Delete(&History{}, "id = ?", id).Error
// }
