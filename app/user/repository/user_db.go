package repository

import "gorm.io/gorm"

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryDB{db: db}
}

func (r *userRepositoryDB) Create(user User) (*User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryDB) GetUserByFirebaseID(firebaseID string) (*User, error) {
	var user User
	if err := r.db.First(&user, "firebase_id = ?", firebaseID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
