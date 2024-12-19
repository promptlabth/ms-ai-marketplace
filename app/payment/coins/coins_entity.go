package coins

import (
	"gorm.io/gorm"
)

type CoinsEntity struct {
	gorm.Model
	FirebaseID string `gorm:"column:firebase_id"`
	AgentID    int    `gorm:"column:agent_id"`
	Coins      int    `gorm:"column:coins"`
}

type CoinReq struct {
	gorm.Model
	AgentID    int    `gorm:"column:agent_id"`
}
