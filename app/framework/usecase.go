package framework

import (
	"context"
	"log"
)

type storage interface {
	CreateFramework(context.Context, FrameworkEntity) (*string, error)
	GetFrameworkByID(context.Context, string) (*FrameworkEntity, error)
	ListFrameworks(context.Context) (*[]FrameworkEntity, error)  
	UpdateFramework(context.Context, FrameworkEntity) error
	DeleteFramework(context.Context, string) error
}

type domain interface {
	ValidateNewFramework(ctx context.Context, framework Framework) error
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

// NewRole orchestrates the process of validating and creating a new role.
func (u *Usecase) NewFramework(ctx context.Context, framework Framework) error {

	// First, use the domain logic to validate the new framework.
	if err := u.domain.ValidateNewFramework(ctx, framework); err != nil {
		log.Printf("Error validating new framework: %v", err)
		return err
	}

	frameworkEntity := FrameworkEntity{
		ID:   framework.ID,
		Name: framework.Name,
		Detail: framework.Detail,
		Component: framework.Component,
	}
	log.Printf("FrameworkEntity: %+v\n", frameworkEntity)

	// If validation passes, proceed to create the frameworkEntity in the storage layer.
	_, err := u.storage.CreateFramework(ctx, frameworkEntity)
	return err
}

func (u *Usecase) ListFrameworks(ctx context.Context) (*[]FrameworkEntity, error) {
	frameworks, err := u.storage.ListFrameworks(ctx)
    if err != nil {
        return nil, err
    }
    return frameworks, nil
}

