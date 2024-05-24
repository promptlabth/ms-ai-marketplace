package roleframework

type RoleFrameworkEntity struct{
	ID int64 `gorm:"autoIncrement;column:id"`
	Name string `gorm:"column:name"`
}

func (RoleFrameworkEntity) TableName() string {
	return "role_frameworks"
}