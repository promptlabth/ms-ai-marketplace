package styleprompt

type StylePromptEntity struct {
    ID   uint   `gorm:"autoIncrement;column:id"`
    Name string `gorm:"column:name"`
    Language  string  `gorm:"column:language"`
}

func (StylePromptEntity) TableName() string {
    return "style_prompts"
}
