package coins

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

func (c *Core) CreateCoins(ctx context.Context, coins CoinsEntity) (*uint, error) {
	if err := c.db.Create(&coins).Error; err != nil {
		return nil, err
	}
	return &coins.ID, nil
}

func (c *Core) GetSumOfCoinsByFirebaseID(ctx context.Context, firebaseID string) (int, error) {
	var totalCoins int
	if err := c.db.Model(&CoinsEntity{}).Where("firebase_id = ?", firebaseID).Select("SUM(coins)").Scan(&totalCoins).Error; err != nil {
		return 0, err
	}
	return totalCoins, nil
}

func (c *Core) GetSumOfCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) (int, error) {
	var totalCoins int
	if err := c.db.Model(&CoinsEntity{}).Where("firebase_id = ? AND agent_id = ?", firebaseID, agentID).Select("SUM(coins)").Scan(&totalCoins).Error; err != nil {
		return 0, err
	}
	return totalCoins, nil
}

func (c *Core) SetCoinsToZeroByFirebaseID(ctx context.Context, firebaseID string) error {
	if err := c.db.Model(&CoinsEntity{}).Where("firebase_id = ?", firebaseID).Update("coins", 0).Error; err != nil {
		return err
	}
	return nil
}

func (c *Core) SetCoinsToZeroByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error {
	if err := c.db.Model(&CoinsEntity{}).Where("firebase_id = ? AND agent_id = ?", firebaseID, agentID).Update("coins", 0).Error; err != nil {
		return err
	}
	return nil
}