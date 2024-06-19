package framework

import "encoding/json"

type FrameworkEntity struct {
	ID        int             `gorm:"autoIncrement;column:id"`
	Name      string          `gorm:"column:name;uniqueIndex"`
	Detail    string          `gorm:"column:detail"`
	Component json.RawMessage `gorm:"column:component"`
	Language  string          `gorm:"column:language"`
}

func (FrameworkEntity) TableName() string {
	return "frameworks"
}
