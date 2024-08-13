package routes

import (
	handlers_book "github.com/aebalz/go-gin-gone/handlers/book"
	"github.com/aebalz/go-gin-gone/models"
	repositories_book "github.com/aebalz/go-gin-gone/repositories/book"
	services_book "github.com/aebalz/go-gin-gone/services/book"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitialRoutesApp(r *gin.Engine, db *gorm.DB) {
	// Database Migrates
	db.AutoMigrate(&models.Author{}, &models.Book{})

	// Repositories
	authorRepo := repositories_book.NewAuthorRepository(db)
	bookRepo := repositories_book.NewBookRepository(db)

	// Services
	authorService := services_book.NewAuthorService(authorRepo)
	bookService := services_book.NewBookService(bookRepo)

	// Handlers

	// Root Routes : "{host}/api/v1/..."
	api := r.Group("/api/v1")
	handlers_book.RegisterAuthorRoutes(api, authorService)
	handlers_book.RegisterBookRoutes(api, bookService)
}
