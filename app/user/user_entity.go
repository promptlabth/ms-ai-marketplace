package user

import "time"

// is a interface to connect to table user on database

type UserEntity struct {
	FirebaseID     string    `gorm:"column:firebase_id"`
	Name           string    `gorm:"column:name"`
	Email          string    `gorm:"column:profile_pic"`
	Platform       string    `gorm:"column:access_token"`
	StripeID       string    `gorm:"column:stripe_id"`
	PlanID         string    `gorm:"column:plan_id"`
	LastActiveTime time.Time `gorm:"column:datetime_last_active"`
}

func (UserEntity) TableName() string {
	return "users"
}
