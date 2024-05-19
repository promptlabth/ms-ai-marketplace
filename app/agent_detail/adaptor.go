// adaptor.go

package agentdetail

import (
	// "context"
	// "errors"
	// "regexp"
	"gorm.io/gorm"
)


type Adaptor struct {
	db *gorm.DB
}

func NewAdaptor(db *gorm.DB) *Adaptor {
	return &Adaptor{db: db}
}