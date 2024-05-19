package framework

import "encoding/json"

type FrameworkEntity struct {
	ID        string          `gorm:"column:id;uniqueIndex"`
	Name      string          `gorm:"column:name;uniqueIndex"`
	Detail    string          `gorm:"column:detail"`
	InputJSON json.RawMessage `gorm:"column:input_json"`
}

func (FrameworkEntity) TableName() string {
	return "frameworks"
}