package coins

import (
	"context"
	
)

type storage interface {
	CreateCoins(ctx context.Context, coins CoinsEntity) (*uint, error)
	GetSumOfCoinsByFirebaseID(ctx context.Context, firebaseID string) (int, error)
	GetSumOfCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) (int, error)
	SetCoinsToZeroByFirebaseID(ctx context.Context, firebaseID string) error
	SetCoinsToZeroByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error
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

func (u *Usecase) GetSumOfCoinsByFirebaseID(ctx context.Context, firebaseID string) (int, error) {
	totalCoins, err := u.storage.GetSumOfCoinsByFirebaseID(ctx, firebaseID)
	if err != nil {
		return 0, err
	}
	return totalCoins, nil
}

func (u *Usecase) GetSumOfCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) (int, error) {
	return u.storage.GetSumOfCoinsByFirebaseIDAndAgentID(ctx, firebaseID, agentID)
}

func (u *Usecase) SetCoinsToZeroByFirebaseID(ctx context.Context, firebaseID string) error {
	err := u.storage.SetCoinsToZeroByFirebaseID(ctx, firebaseID)
	return err
}

func (u *Usecase) SetCoinsToZeroByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error {
	return u.storage.SetCoinsToZeroByFirebaseIDAndAgentID(ctx, firebaseID, agentID)
}