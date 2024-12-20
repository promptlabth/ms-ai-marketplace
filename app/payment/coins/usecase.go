package coins

import (
	"context"
	
)

type storage interface {
	CreateCoins(ctx context.Context, coins CoinsEntity) (*uint, error)
	GetSumOfCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) (int, error)
	SetCoinsToZeroByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error
	IncreaseCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error
}

type domain interface {
	ValidateNewCoins(ctx context.Context, coins CoinsEntity) error
}

type Usecase struct {
	storage storage
	domain  domain
}

func NewUsecase(s storage, d domain) *Usecase {
	return &Usecase{
		storage: s,
		domain:  d,
	}
}

func (u *Usecase) CreateCoins(ctx context.Context, coins CoinsEntity) (*uint, error) {
	err := u.domain.ValidateNewCoins(ctx, coins)
	if err != nil {
		return nil, err
	}

	return u.storage.CreateCoins(ctx, coins)
}

func (u *Usecase) AddCoins(ctx context.Context, coins CoinsEntity) error {
	existingCoins, err := u.storage.GetSumOfCoinsByFirebaseIDAndAgentID(ctx, coins.FirebaseID, coins.AgentID)
	if err != nil {
		return err
	}

	coins.Coins += existingCoins

	_, err = u.storage.CreateCoins(ctx, coins)
	return err
}

func (u *Usecase) GetSumOfCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) (int, error) {
	return u.storage.GetSumOfCoinsByFirebaseIDAndAgentID(ctx, firebaseID, agentID)
}

func (u *Usecase) SetCoinsToZeroByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error {
	return u.storage.SetCoinsToZeroByFirebaseIDAndAgentID(ctx, firebaseID, agentID)
}

// IncreaseCoinsByFirebaseIDAndAgentID increases the coins by 1 for a specific firebase_id and agent_id
func (u *Usecase) IncreaseCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error {
	return u.storage.IncreaseCoinsByFirebaseIDAndAgentID(ctx, firebaseID, agentID)
}