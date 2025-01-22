package repository

type User struct {
	ID                 int    `db:"id"`
	FirebaseID         string `db:"firebase_id"`
	Name               string `db:"name"`
	Email              string `db:"email"`
	Platform           string `db:"platform"`
	StripeID           string `db:"stripe_id"`
	PlanID             string `db:"plan_id"`
	DatetimeLastActive string `db:"datetime_last_active"`
	ProfilePicture     string `db:"profile_pic"`
	AccessToken        string `db:"access_token"`
	Role               string `db:"role"`
	MaxMessages        int    `db:"max_messages"`
}

type UserRepository interface {
	Create(User) (*User, error)
	GetUserByFirebaseID(string) (*User, error)
	Update(User) (*User, error)
}
