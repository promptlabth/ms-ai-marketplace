package realtimegen

import (
	// "context"
	"gorm.io/gorm"
)

// Adaptor is responsible for the data storage operations and validation logic related to histories.
type Adaptor struct {
	db *gorm.DB
}

// NewAdaptor creates a new instance of Adaptor with a database connection.
func NewAdaptor(db *gorm.DB) *Adaptor {
	return &Adaptor{db: db}
}

