package framework

import (
	"context"
	"log"
)

type storage interface {
	CreateFramework(context.Context, FrameworkEntity) (*string, error)
	GetFrameworkByID(context.Context, int) (*FrameworkEntity, error)
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

	frameworkEntity := FrameworkEntity{
		Name:      framework.Name,
		Detail:    framework.Detail,
		Component: framework.Component,
	}
	log.Printf("FrameworkEntity: %+v\n", frameworkEntity)

	_, err := u.storage.CreateFramework(ctx, frameworkEntity)
	return err
}



func (u *Usecase) GetFrameworkByID(ctx context.Context, id int) (*FrameworkEntity, error) {
	framework, err := u.storage.GetFrameworkByID(ctx, id)
	if err != nil {
		log.Printf("Error getting framwork by ID: %v", err)
		return nil, err
	}
	return framework, nil
}

func (u *Usecase) ListFrameworks(ctx context.Context) (*[]FrameworkEntity, error) {
	frameworks, err := u.storage.ListFrameworks(ctx)
	if err != nil {
		return nil, err
	}
	return frameworks, nil
}
