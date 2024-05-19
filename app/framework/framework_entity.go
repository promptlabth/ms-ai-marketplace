package framework

type FrameworkEntity struct{
	Name string `gorm:"column:name"`
}

func (FrameworkEntity) TableName() string {
	return "frameworks"
}