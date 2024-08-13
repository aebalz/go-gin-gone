package handlers_book

import (
	services_book "github.com/aebalz/go-gin-gone/services/book"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.RouterGroup, bookService services_book.BookService) {
	books := r.Group("/books")
	{
		books.GET("/", bookService.GetBooks)
		books.GET("/:id", bookService.GetBook)
		books.POST("/", bookService.CreateBook)
		books.PUT("/:id", bookService.UpdateBook)
		books.PATCH("/:id", bookService.PatchBook)
		books.DELETE("/:id", bookService.DeleteBook)
	}
}
