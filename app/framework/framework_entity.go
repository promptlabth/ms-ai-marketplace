package framework

import "encoding/json"

type FrameworkEntity struct {
	ID        int64          `gorm:"autoIncrement;column:id "`
	Name      string          `gorm:"column:name;uniqueIndex"`
	Detail    string          `gorm:"column:detail"`
	Component json.RawMessage `gorm:"column:component"`
}

func (FrameworkEntity) TableName() string {
	return "frameworks"
}