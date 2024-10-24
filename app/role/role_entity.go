package role

// RoleEntity is an interface to connect to the 'roles' table in the database
type RoleEntity struct {
    ID   int   `gorm:"autoIncrement;column:id"`
    Name string `gorm:"column:name"`
    Language  string  `gorm:"column:language"`
}

// TableName sets the insert table name for this struct type
func (RoleEntity) TableName() string {
    return "roles"
}
