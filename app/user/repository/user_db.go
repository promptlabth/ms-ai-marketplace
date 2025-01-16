package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

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
    // log.Printf("FirebaseID in service: %s", firebaseID)
    // fmt.Println("FirebaseID in service: %s", firebaseID)
    if firebaseID == "" {
        return nil, errors.New("firebaseID cannot be empty")
    }

    var user User
    if err := r.db.First(&user, "firebase_id = ?", firebaseID).Error; err != nil {
        log.Printf("Error querying user: %v", err)
        return nil, err
    }
    return &user, nil
}

func (r *userRepositoryDB) Update(user User) (*User, error) {
    if err := r.db.Save(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}