package repositories_book

import (
	"github.com/aebalz/go-gin-gone/models"
	"github.com/aebalz/go-gin-gone/utils"
	"github.com/aebalz/go-gin-gone/utils/paginate"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll(p *paginate.Param) ([]models.Author, int64, error)
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

func (r *authorRepository) FindAll(p *paginate.Param) ([]models.Author, int64, error) {
	var authors []models.Author
	var count int64
	result := r.db.Model(&models.Author{}).Preload("Books").Count(&count).Scopes(paginate.GormPaginate(p)).Find(&authors)
	return authors, count, result.Error
}

func (r *authorRepository) FindByID(id uint) (models.Author, error) {
	var author models.Author
	result := r.db.Model(&models.Author{}).First(&author, id)
	return author, result.Error
}

func (r *authorRepository) Create(author models.Author) (models.Author, error) {
	// Validate the author before creating
	if err := utils.ValidateStruct(author); err != nil {
		return author, err
	}

	result := r.db.Model(&models.Author{}).Create(&author)
	return author, result.Error
}

func (r *authorRepository) Update(author models.Author) (models.Author, error) {
	// Validate the author before creating
	if err := utils.ValidateStruct(author); err != nil {
		return author, err
	}

	result := r.db.Model(&models.Author{}).Save(&author)
	return author, result.Error
}

func (r *authorRepository) Delete(id uint) error {
	result := r.db.Model(&models.Author{}).Delete(&models.Author{}, id)
	return result.Error
}
