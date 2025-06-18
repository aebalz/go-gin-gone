package handlers_book

import (
	"net/http"
	"strconv"

	"github.com/aebalz/go-gin-gone/models"
	services_book "github.com/aebalz/go-gin-gone/services/book"
	"github.com/aebalz/go-gin-gone/utils/paginate" // For GetPaginateParam and PaginateRes
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorService services_book.AuthorService
}

func NewAuthorHandler(authorService services_book.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

// RegisterAuthorRoutes needs to be defined to register routes for author actions.
// This was mentioned in the main task to be updated.
func RegisterAuthorRoutes(r *gin.RouterGroup, authorHandler *AuthorHandler) {
	authors := r.Group("/authors")
	{
		authors.GET("/", authorHandler.GetAuthors)
		authors.GET("/:id", authorHandler.GetAuthor)
		authors.POST("/", authorHandler.CreateAuthor)
		authors.PUT("/:id", authorHandler.UpdateAuthor)
		authors.DELETE("/:id", authorHandler.DeleteAuthor)
	}
}

func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	p := paginate.GetPaginateParam(c)
	authors, count, err := h.authorService.GetAuthors(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, paginate.PaginateRes[[]models.Author]{
		Data: authors,
		Paginate: paginate.PaginateMeta{
			LastPage:    paginate.CalculateLastPage(count, p.Limit),
			CurrentPage: p.Page,
			Limit:       p.Limit,
			Total:       count,
		},
	})
}

func (h *AuthorHandler) GetAuthor(c *gin.Context) {
	idStr := c.Param("id")
	authorID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	author, err := h.authorService.GetAuthor(uint(authorID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var dto services_book.CreateAuthorDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, err := h.authorService.CreateAuthor(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, author)
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	idStr := c.Param("id")
	authorID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var dto services_book.UpdateAuthorDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := h.authorService.UpdateAuthor(uint(authorID), dto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found or could not be updated"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")
	authorID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.authorService.DeleteAuthor(uint(authorID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found or could not be deleted"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
