// user.go

package user

import (
	"context"
	"time"
)

// User represents the structure of a user in the system.
type User struct {
	FriebaseID     string
	Name           string
	Email          string
	Platform       string // Note: In a real system, you would not store plain text passwords
	StripeID       string
	PlanID         string
	Password       string
	LastActiveTime time.Time
}

// UserInterface defines the set of methods that any implementation of the User service must provide.
type UserInterface interface {
	CreateUser(ctx context.Context, user User) (int64, error) // Creates a new user and returns the user ID
	GetUserByID(ctx context.Context, id int64) (*User, error) // Fetches a user by their ID
	UpdateUser(ctx context.Context, user User) error          // Updates an existing user
	DeleteUser(ctx context.Context, id int64) error           // Deletes a user by their ID
}

type NewUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Include other fields as necessary
}
