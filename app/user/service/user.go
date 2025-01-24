package service

import "context"

type NewUserRequest struct {
	FirebaseID         string `json:"firebase_id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Platform           string `json:"platform"`
	StripeID           string `json:"stripe_id"`
	PlanID             string `json:"plan_id"`
	DatetimeLastActive string `json:"datetime_last_active"`
	ProfilePicture     string `json:"profile_pic"`
	AccessToken        string `json:"access_token"`
}

type UserResponse struct {
	ID             int    `json:"id"`
	FirbaseID      string `json:"firebase_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Platform       string `json:"platform"`
	PlanID         string `json:"plan_id"`
	ProfilePicture string `json:"profile_pic"`
	AccessToken    string `json:"access_token"`
	Role           string `json:"role"`
	MaxMessages    int    `json:"max_messages"`
	UsedMessages   int    `json:"used_messages"`
}

type UserService interface {
	NewUser(ctx context.Context, request NewUserRequest) (*UserResponse, error)
	GetUser(string) (UserResponse, error)
}

type PromplabResponse struct {
	User User `json:"user"`
	Plan Plan `json:"plan"`
}

type Product struct {
	ID          int    `json:"id"`
	PlanType    string `json:"planType"`
	MaxMessages int    `json:"maxMessages"`
	ProductID   *int   `json:"product_id"`
}

type Plan struct {
	Product   Product `json:"product"`
	StartDate *string `json:"start_date"`
	EndDate   *string `json:"end_date"`
}

type User struct {
	ID          int     `json:"id"`
	FirebaseID  string  `json:"firebase_id"`
	Name        string  `json:"name"`
	Email       *string `json:"email"`
	ProfilePic  string  `json:"profilepic"`
	Platform    string  `json:"platform"`
	AccessToken string  `json:"access_token"`
	StripeID    string  `json:"stripe_id"`
	PlanID      int     `json:"plan_id"`
}