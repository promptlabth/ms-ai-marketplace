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


func (c *Core) GetHistoryByFirebaseID(ctx context.Context, firebaseID string) ([]HistoryWithAgentDetail, error) {
    var histories []HistoryWithAgentDetail
    subQuery := c.db.Table("histories").
        Select("MAX(time_stamp)").
        Where("firebase_id = ?", firebaseID).
        Group("agent_id")

    if err := c.db.Table("histories").
        Select("histories.*, agent_details.name, agent_details.description, agent_details.image_url, agent_details.prompt, agent_details.framework_id, agent_details.role_framework_id, agent_details.total_used").
        Joins("left join agent_details on agent_details.id = histories.agent_id").
        Where("histories.time_stamp IN (?)", subQuery).
        Order("histories.time_stamp DESC").
        Scan(&histories).Error; err != nil {
        return nil, err
    }
    return histories, nil
}

func (c *Core) GetHistoriesByAgentIDs(ctx context.Context, agentIDs []int) ([]HistoryWithAgentDetail, error) {
    var histories []HistoryWithAgentDetail
    if err := c.db.Table("histories").
        Select("histories.*, agent_details.name, agent_details.description, agent_details.image_url, agent_details.prompt, agent_details.framework_id, agent_details.role_framework_id, agent_details.total_used").
        Joins("left join agent_details on agent_details.id = histories.agent_id").
        Where("histories.agent_id IN ?", agentIDs).
        Order("histories.time_stamp DESC").
        Scan(&histories).Error; err != nil {
        return nil, err
    }
    return histories, nil
}
