package roleframework

type RoleFrameworkEntity struct{
	Name string `gorm:"column:name"`
}

func (RoleFrameworkEntity) TableName() string {
	return "role_frameworks"
}