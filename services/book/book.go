package services_book

// CreateBookDTO defines the data transfer object for creating a new book.
type CreateBookDTO struct {
	Title    string  `json:"title"`
	AuthorID uint    `json:"author_id"`
	ISBN     string  `json:"isbn"`
	Price    float64 `json:"price"`
}

// UpdateBookDTO defines the data transfer object for updating an existing book.
type UpdateBookDTO struct {
	Title    string  `json:"title"`
	AuthorID uint    `json:"author_id"`
	ISBN     string  `json:"isbn"`
	Price    float64 `json:"price"`
}

// PatchBookDTO defines the data transfer object for partially updating an existing book.
type PatchBookDTO struct {
	Title    *string  `json:"title"`
	AuthorID *uint    `json:"author_id"`
	ISBN     *string  `json:"isbn"`
	Price    *float64 `json:"price"`
}

import (
	"github.com/aebalz/go-gin-gone/models"
	"github.com/aebalz/go-gin-gone/utils/paginate"
)

// BookRepository defines the methods that any
// data storage provider needs to implement to get
// and store books
type BookRepository interface {
	FindAll(p *paginate.Param) ([]models.Book, int64, error)
	FindByID(id uint) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(id uint) error
}

// BookService defines the methods that a book service should implement
type BookService interface {
	GetBooks(p *paginate.Param) ([]models.Book, int64, error)
	GetBook(id uint) (models.Book, error)
	CreateBook(bookDTO CreateBookDTO) (models.Book, error)
	UpdateBook(id uint, bookDTO UpdateBookDTO) (models.Book, error)
	PatchBook(id uint, bookDTO PatchBookDTO) (models.Book, error)
	DeleteBook(id uint) error
}

// bookService implements the BookService interface
type bookService struct {
	repo BookRepository
}

// NewBookService creates a new book service
func NewBookService(repo BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetBooks(p *paginate.Param) ([]models.Book, int64, error) {
	books, count, err := s.repo.FindAll(p)
	if err != nil {
		return nil, 0, err
	}
	return books, count, nil
}

func (s *bookService) GetBook(id uint) (models.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s *bookService) CreateBook(bookDTO CreateBookDTO) (models.Book, error) {
	newBook := models.Book{
		Title:    bookDTO.Title,
		AuthorID: bookDTO.AuthorID,
		ISBN:     bookDTO.ISBN,
		Price:    bookDTO.Price,
	}
	book, err := s.repo.Create(newBook)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s *bookService) UpdateBook(id uint, bookDTO UpdateBookDTO) (models.Book, error) {
	updatedBook := models.Book{
		ID:       id,
		Title:    bookDTO.Title,
		AuthorID: bookDTO.AuthorID,
		ISBN:     bookDTO.ISBN,
		Price:    bookDTO.Price,
	}
	book, err := s.repo.Update(updatedBook)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s *bookService) PatchBook(id uint, bookDTO PatchBookDTO) (models.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return models.Book{}, err
	}

	if bookDTO.Title != nil {
		book.Title = *bookDTO.Title
	}
	if bookDTO.AuthorID != nil {
		book.AuthorID = *bookDTO.AuthorID
	}
	if bookDTO.ISBN != nil {
		book.ISBN = *bookDTO.ISBN
	}
	if bookDTO.Price != nil {
		book.Price = *bookDTO.Price
	}

	updatedBook, err := s.repo.Update(book)
	if err != nil {
		return models.Book{}, err
	}
	return updatedBook, nil
}

func (s *bookService) DeleteBook(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
