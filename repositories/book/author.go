package repositories_book

import (
	"github.com/aebalz/go-gin-gone/models"
	"github.com/aebalz/go-gin-gone/utils"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() ([]models.Author, error)
	FindByID(id uint) (models.Author, error)
	Create(author models.Author) (models.Author, error)
	Update(author models.Author) (models.Author, error)
	Delete(id uint) error
}

// authorRepository implements the AuthorRepository interface
type authorRepository struct {
	db *gorm.DB
}

// NewAuthorRepository creates a new author repository
func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) FindAll() ([]models.Author, error) {
	var authors []models.Author
	result := r.db.Preload("Books").Find(&authors)
	return authors, result.Error
}

func (r *authorRepository) FindByID(id uint) (models.Author, error) {
	var author models.Author
	result := r.db.First(&author, id)
	return author, result.Error
}

func (r *authorRepository) Create(author models.Author) (models.Author, error) {
	// Validate the author before creating
	if err := utils.ValidateStruct(author); err != nil {
		return author, err
	}

	result := r.db.Create(&author)
	return author, result.Error
}

func (r *authorRepository) Update(author models.Author) (models.Author, error) {
	// Validate the author before creating
	if err := utils.ValidateStruct(author); err != nil {
		return author, err
	}

	result := r.db.Save(&author)
	return author, result.Error
}

func (r *authorRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Author{}, id)
	return result.Error
}
