package models

type Author struct {
	BaseModel
	Name  string `gorm:"size:255;not null" validate:"required" json:"name"`
	Books []Book `gorm:"foreignKey:AuthorID" validate:"dive" json:"books"`
}

// Implement the TableName method for Author
func (Author) TableName() string {
	return apisPrefixName + "author"
}
