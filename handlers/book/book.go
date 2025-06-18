package handlers_book

import (
	"net/http"
	"strconv"

	"github.com/aebalz/go-gin-gone/models" // Keep for response if service returns models.Book
	services_book "github.com/aebalz/go-gin-gone/services/book"
	"github.com/aebalz/go-gin-gone/utils/paginate" // For GetPaginateParam and PaginateRes
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService services_book.BookService
}

func NewBookHandler(bookService services_book.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func RegisterBookRoutes(r *gin.RouterGroup, bookHandler *BookHandler) {
	books := r.Group("/books")
	{
		books.GET("/", bookHandler.GetBooks)
		books.GET("/:id", bookHandler.GetBook)
		books.POST("/", bookHandler.CreateBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.PATCH("/:id", bookHandler.PatchBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	p := paginate.GetPaginateParam(c)
	books, count, err := h.bookService.GetBooks(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, paginate.PaginateRes[[]models.Book]{
		Data: books,
		Paginate: paginate.PaginateMeta{
			LastPage:    paginate.CalculateLastPage(count, p.Limit),
			CurrentPage: p.Page,
			Limit:       p.Limit,
			Total:       count,
		},
	})
}

func (h *BookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	bookID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	book, err := h.bookService.GetBook(uint(bookID))
	if err != nil {
		// This could be a StatusNotFound or StatusInternalServerError depending on the error
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"}) // Assuming service returns specific error for not found
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) PatchBook(c *gin.Context) {
	idStr := c.Param("id")
	bookID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dto services_book.PatchBookDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.PatchBook(uint(bookID), dto)
	if err != nil {
		// Could be StatusNotFound or StatusInternalServerError
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or could not be patched"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var dto services_book.CreateBookDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := h.bookService.CreateBook(dto)
	if err != nil {
		// Depending on the error type, could be StatusBadRequest or StatusInternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	bookID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.bookService.DeleteBook(uint(bookID))
	if err != nil {
		// Could be StatusNotFound or StatusInternalServerError
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or could not be deleted"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	bookID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dto services_book.UpdateBookDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.UpdateBook(uint(bookID), dto)
	if err != nil {
		// Could be StatusNotFound or StatusInternalServerError
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or could not be updated"})
		return
	}
	c.JSON(http.StatusOK, book)
}
