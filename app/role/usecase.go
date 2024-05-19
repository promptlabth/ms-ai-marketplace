package role

import (
	"context"
	"log"
)

// Assuming Role struct is defined in role.go within the 'user' package.

// storage outlines the methods required by the use case to interact with the data layer.
type storage interface {
	CreateRole(ctx context.Context, role RoleEntity) (*int, error)
	GetRoleByID(ctx context.Context, id int) (*RoleEntity, error)
	UpdateRole(ctx context.Context, role RoleEntity) error
	DeleteRole(ctx context.Context, id int) error
	// Add other storage methods as necessary
}

// domain outlines the methods required by the use case for domain logic and validations.
type domain interface {
	ValidateNewRole(ctx context.Context, role Role) error
	// Add other domain methods as necessary
}

// Usecase struct that combines storage and domain to execute role-related business logic.
type Usecase struct {
	storage storage
	domain  domain
}

// NewUsecase creates a new Usecase instance with the provided storage and domain logic implementations.
func NewUsecase(s storage, d domain) *Usecase {
	return &Usecase{
		storage: s,
		domain:  d,
	}
}

// NewRole orchestrates the process of validating and creating a new role.
func (u *Usecase) NewRole(ctx context.Context, role Role) error {

	// First, use the domain logic to validate the new role.
	if err := u.domain.ValidateNewRole(ctx, role); err != nil {
		log.Printf("Error validating new role: %v", err)
		return err
	}

	roleEntity := RoleEntity{
		ID:   role.ID,
		Name: role.Name,
	}
	log.Printf("RoleEntity: %+v\n", roleEntity)

	// If validation passes, proceed to create the role in the storage layer.
	_, err := u.storage.CreateRole(ctx, roleEntity)
	return err
}
