package services_book

import (
	"net/http"
	"strconv"

	"github.com/aebalz/go-gin-gone/models"
	repositories_book "github.com/aebalz/go-gin-gone/repositories/book"
	"github.com/gin-gonic/gin"
)

// AuthorService defines the methods that a author service should implement
type AuthorService interface {
	GetAuthors(c *gin.Context)
	GetAuthor(c *gin.Context)
	CreateAuthor(c *gin.Context)
	UpdateAuthor(c *gin.Context)
	DeleteAuthor(c *gin.Context)
}

// authorService implements the AuthorService interface
type authorService struct {
	repo repositories_book.AuthorRepository
}

// NewAuthorService creates a new author service
func NewAuthorService(repo repositories_book.AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) GetAuthors(c *gin.Context) {
	authors, err := s.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authors)
}

func (s *authorService) GetAuthor(c *gin.Context) {
	id := c.Param("id")
	authorID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	author, err := s.repo.FindByID(uint(authorID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func (s *authorService) CreateAuthor(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, err := s.repo.Create(newAuthor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, author)
}

func (s *authorService) UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var updatedAuthor models.Author
	if err := c.ShouldBindJSON(&updatedAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authorID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	updatedAuthor.ID = uint(authorID)
	author, err := s.repo.Update(updatedAuthor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)
}

func (s *authorService) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	authorID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = s.repo.Delete(uint(authorID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
