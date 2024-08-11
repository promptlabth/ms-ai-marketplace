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
	ListRoles(ctx context.Context,language string) (*[]RoleEntity, error)
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
		Name: role.Name,
	}
	log.Printf("RoleEntity: %+v\n", roleEntity)

	// If validation passes, proceed to create the role in the storage layer.
	_, err := u.storage.CreateRole(ctx, roleEntity)
	return err
}

func (u *Usecase) GetRoleByID(ctx context.Context, id int) (*RoleEntity, error) {
    role, err := u.storage.GetRoleByID(ctx, id)
    if err != nil {
        log.Printf("Error getting role by ID: %v", err)
        return nil, err
    }
    return role, nil
}

func (u *Usecase) UpdateRole(ctx context.Context, role RoleEntity) error {
    err := u.storage.UpdateRole(ctx, role)
    if err != nil {
        log.Printf("Error updating role: %v", err)
        return err
    }
    return nil
}


func (u *Usecase) DeleteRole(ctx context.Context, id int) error {
    err := u.storage.DeleteRole(ctx, id)
    if err != nil {
        log.Printf("Error deleting role: %v", err)
        return err
    }
    return nil
}


// Get List of Roles
func (u *Usecase) ListRoles(ctx context.Context,language string) (*[]RoleEntity, error) {
	roles, err := u.storage.ListRoles(ctx,language)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
