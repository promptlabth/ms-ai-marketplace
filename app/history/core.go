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


func (c *Core) GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryEntity, error) {
    var histories []HistoryEntity
    subQuery := c.db.Table("histories").
        Select("MAX(time_stamp)").
        Where("firebase_id = ?", firebaseID).
        Group("agent_id")

    if err := c.db.Where("time_stamp IN (?)", subQuery).Order("time_stamp DESC").Find(&histories).Error; err != nil {
        return nil, err
    }
    return histories, nil
}