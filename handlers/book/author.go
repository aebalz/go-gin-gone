package handlers_book

import (
	services_book "github.com/aebalz/go-gin-gone/services/book"
	"github.com/gin-gonic/gin"
)

func RegisterAuthorRoutes(r *gin.RouterGroup, authorService services_book.AuthorService) {
	authors := r.Group("/authors")
	{
		authors.GET("/", authorService.GetAuthors)
		authors.GET("/:id", authorService.GetAuthor)
		authors.POST("/", authorService.CreateAuthor)
		authors.PUT("/:id", authorService.UpdateAuthor)
		authors.PATCH("/:id", authorService.UpdateAuthor)
		authors.DELETE("/:id", authorService.DeleteAuthor)
	}
}
