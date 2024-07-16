// adaptor.go

package history

import (
	"context"
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
	// if history.UserID == "" {
	// 	return errors.New("User ID cannot be zero")
	// }

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

// func (a *Adaptor) existsInDatabase(ctx context.Context, tableName string, id string) bool {
// 	var exists bool
// 	query := "SELECT EXISTS (SELECT 1 FROM " + tableName + " WHERE id = ?)"
// 	err := a.db.Raw(query, id).Scan(&exists).Error
// 	if err != nil {
// 		log.Printf("Error checking existence in %s: %v", tableName, err)
// 		return false
// 	}
// 	return exists
// }
