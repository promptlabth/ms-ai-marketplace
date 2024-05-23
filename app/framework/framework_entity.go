package framework

import "encoding/json"

type FrameworkEntity struct {
	ID        string          `gorm:"column:id; uniqueIndex "`
	Name      string          `gorm:"column:name;uniqueIndex"`
	Detail    string          `gorm:"column:detail"`
	Component json.RawMessage `gorm:"column:component"`
}

func (FrameworkEntity) TableName() string {
	return "frameworks"
}