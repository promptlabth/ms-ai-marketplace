// core.go

package user

import (
	"context"
	"database/sql"
)

// Core is responsible for the data storage operations related to users.
type Core struct {
	db *sql.DB
}

// NewCore creates a new instance of Core with a database connection.
func NewCore(db *sql.DB) *Core {
	return &Core{db: db}
}

// CreateUser inserts a new user into the database.
func (c *Core) CreateUser(ctx context.Context, user User) (int64, error) {
	const query = `
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?)
		RETURNING id
	`

	// Assuming `User` struct has `Username`, `Email`, and `Password` fields.
	// Adjust the query placeholders and fields based on your actual database and `User` struct.
	var userID int64
	err := c.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// GetUserByID retrieves a user by their ID from the database.
func (c *Core) GetUserByID(ctx context.Context, id int64) (*User, error) {
	const query = `
		SELECT id, username, email, password
		FROM users
		WHERE id = ?
	`

	var user User
	err := c.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user's information in the database.
func (c *Core) UpdateUser(ctx context.Context, user User) error {
	const query = `
		UPDATE users
		SET username = ?, email = ?, password = ?
		WHERE id = ?
	`

	_, err := c.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.ID)
	return err
}

// DeleteUser removes a user from the database by their ID.
func (c *Core) DeleteUser(ctx context.Context, id int64) error {
	const query = `
		DELETE FROM users
		WHERE id = ?
	`

	_, err := c.db.ExecContext(ctx, query, id)
	return err
}
