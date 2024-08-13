package models

type Book struct {
	BaseModel
	Title    string  `gorm:"size:255;not null" validate:"required" json:"title"`
	AuthorID uint    `gorm:"not null" json:"author_id"`
	ISBN     string  `gorm:"size:20" validate:"required,len=13,numeric" json:"isbn"`
	Price    float64 `gorm:"type:decimal(10,2)" validate:"required,gt=0" json:"price"`
	// Publisher   string  `gorm:"size:255" validate:"required" json:"publisher"`
	// PublishedAt string  `gorm:"size:255" validate:"required,datetime" json:"published_at"`
	// Author      Author  `gorm:"foreignKey:AuthorID"`
}

func (Book) TableName() string {
	return apisPrefixName + "book"
}
