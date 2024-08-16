// user.go

package user

import (
	"context"
	"time"
)

// User represents the structure of a user in the system.
type User struct {
	FirebaseID     string
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
	FirebaseID string `json:"firebase_id"`
	Name       string `json:"name"`
	Email      string `json:"profile_pic"`
	Platform   string `json:"access_token"`
	StripeID   string `json:"stripe_id"`
	PlanID     string `json:"plan_id"`
	Password   string `json:"password"`
	// none LastActiveTime time.Time `json:"datetime_last_active"` because of is now().time
	// Include other fields as necessary
}
