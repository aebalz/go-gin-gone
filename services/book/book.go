package services_book

import (
	"net/http"
	"strconv"

	"github.com/aebalz/go-gin-gone/models"
	repositories_book "github.com/aebalz/go-gin-gone/repositories/book"
	"github.com/gin-gonic/gin"
)

// BookService defines the methods that a book service should implement
type BookService interface {
	GetBooks(c *gin.Context)
	GetBook(c *gin.Context)
	CreateBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	PatchBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

// bookService implements the BookService interface
type bookService struct {
	repo repositories_book.BookRepository
}

// NewBookService creates a new book service
func NewBookService(repo repositories_book.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetBooks(c *gin.Context) {
	books, err := s.repo.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (s *bookService) GetBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	book, err := s.repo.FindByID(uint(bookID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (s *bookService) CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := s.repo.Create(newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (s *bookService) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	updatedBook.ID = uint(bookID)
	book, err := s.repo.Update(updatedBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (s *bookService) PatchBook(c *gin.Context) {
	id := c.Param("id")
	var bookUpdates map[string]interface{}
	if err := c.ShouldBindJSON(&bookUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	book, err := s.repo.FindByID(uint(bookID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if title, exists := bookUpdates["title"]; exists {
		book.Title = title.(string)
	}
	// if author, exists := bookUpdates["author"]; exists {
	// 	book.Author = author.(string)
	// }

	if isbn, exists := bookUpdates["isbn"]; exists {
		book.ISBN = isbn.(string)
	}
	if price, exists := bookUpdates["price"]; exists {
		book.Price = price.(float64)
	}
	updatedBook, err := s.repo.Update(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

func (s *bookService) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = s.repo.Delete(uint(bookID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
