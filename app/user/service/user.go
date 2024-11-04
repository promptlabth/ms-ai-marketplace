package service

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
}

type UserService interface {
	NewUser(NewUserRequest) (*UserResponse, error)
	GetUser(string) (UserResponse, error)
}
