package user

import "database/sql"

type Adaptor struct {
	db *sql.DB
}

func NewAdaptor(db *sql.DB) *Adaptor {
	return &Adaptor{db: db}
}

func (s *Adaptor) NewUser() error {
	return nil
}
