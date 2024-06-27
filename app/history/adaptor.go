// adaptor.go

package history

import (
	"context"
	"errors"
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

// ValidateNewHistory performs validation on a new history record.
func (a *Adaptor) ValidateNewHistory(ctx context.Context, history History) error {
	// Validate the user ID is not zero.
	if history.UserID == 0 {
		return errors.New("User ID cannot be zero")
	}

	// // Validate the agent ID is not zero.
	// if history.AgentID == 0 {
	// 	return errors.New("Agent ID cannot be zero")
	// }

	// // Validate the framework ID is not zero.
	// if history.FrameworkID == 0 {
	// 	return errors.New("Framework ID cannot be zero")
	// }

	// // Validate the prompt is not empty.
	// if history.Prompt == "" {
	// 	return errors.New("Prompt cannot be empty")
	// }

	// // Validate the result is not empty.
	// if history.Result == "" {
	// 	return errors.New("Result cannot be empty")
	// }

	// Add additional validations as needed.

	// If all validations pass, no error is returned.
	return nil
}

// CreateHistory inserts a new history record into the database.
func (a *Adaptor) CreateHistory(ctx context.Context, history History) (*int, error) {
	if err := a.db.Create(&history).Error; err != nil {
		return nil, err
	}
	return &history.ID, nil
}

// GetHistoryByID retrieves a history record by its ID from the database.
func (a *Adaptor) GetHistoryByID(ctx context.Context, id int) (*History, error) {
	var history History
	if err := a.db.First(&history, id).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

// ListHistories retrieves history records by user ID from the database.
func (a *Adaptor) ListHistories(ctx context.Context, userID int) (*[]History, error) {
	var histories []History
	query := a.db.WithContext(ctx).Where("user_id = ?", userID).Find(&histories)
	if query.Error != nil {
		return nil, query.Error
	}
	return &histories, nil
}

// UpdateHistory updates a history record's information in the database.
func (a *Adaptor) UpdateHistory(ctx context.Context, history History) error {
	return a.db.Model(&History{}).Where("id = ?", history.ID).Updates(history).Error
}

// DeleteHistory removes a history record from the database by its ID.
func (a *Adaptor) DeleteHistory(ctx context.Context, id int) error {
	return a.db.Delete(&History{}, "id = ?", id).Error
}
