// core.go

package user

import (
	"context"

	"gorm.io/gorm"
)

// Core is responsible for the data storage operations related to users.
type Core struct {
	db *gorm.DB
}

// NewCore creates a new instance of Core with a database connection.
func NewCore(db *gorm.DB) *Core {
	return &Core{db: db}
}

// CreateUser inserts a new user into the database.
func (c *Core) CreateUser(ctx context.Context, user UserEntity) (*string, error) {

	// Assuming `User` struct has `Username`, `Email`, and `Password` fields.
	// Adjust the query placeholders and fields based on your actual database and `User` struct.

	if err := c.db.Create(&user); err.Error != nil {
		return nil, err.Error
	}

	return &user.StripeID, nil
}

// GetUserByID retrieves a user by their ID from the database.
func (c *Core) GetUserByID(ctx context.Context, firebaseId string) (*UserEntity, error) {

	var user UserEntity
	if err := c.db.Where("id = ?", firebaseId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user's information in the database.
func (c *Core) UpdateUser(ctx context.Context, user UserEntity) error {
	return c.db.Model(UserEntity{}).Where("firebase_id = ?", user.FirebaseID).Updates(user).Error
}

// DeleteUser removes a user from the database by their ID.
func (c *Core) DeleteUser(ctx context.Context, id string) error {
	return c.db.Delete(&UserEntity{}, id).Error
}
