package user

import "database/sql"

type Core struct {
	db *sql.DB
}

func NewCore(db *sql.DB) *Core {
	return &Core{db: db}
}
