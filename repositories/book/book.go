package repositories_book

import (
	"github.com/aebalz/go-gin-gone/models"
	"github.com/aebalz/go-gin-gone/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BookRepository defines the methods that any
// data storage provider needs to implement to get
// and store books
type BookRepository interface {
	FindAll(c *gin.Context) ([]models.Book, error)
	FindByID(id uint) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(id uint) error
}

// bookRepository implements the BookRepository interface
type bookRepository struct {
	db *gorm.DB
}

// NewBookRepository creates a new book repository
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll(c *gin.Context) ([]models.Book, error) {
	var books []models.Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *bookRepository) FindByID(id uint) (models.Book, error) {
	var book models.Book
	result := r.db.First(&book, id)
	return book, result.Error
}

func (r *bookRepository) Create(book models.Book) (models.Book, error) {
	if err := utils.ValidateStruct(book); err != nil {
		return book, err
	}

	result := r.db.Create(&book)
	return book, result.Error
}

func (r *bookRepository) Update(book models.Book) (models.Book, error) {
	if err := utils.ValidateStruct(book); err != nil {
		return book, err
	}

	result := r.db.Save(&book)
	return book, result.Error
}

func (r *bookRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Book{}, id)
	return result.Error
}
