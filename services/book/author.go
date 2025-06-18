package services_book

import (
	"github.com/aebalz/go-gin-gone/models"
	"github.com/aebalz/go-gin-gone/utils/paginate"
)

// CreateAuthorDTO defines the data transfer object for creating a new author.
type CreateAuthorDTO struct {
	Name string `json:"name" validate:"required"`
}

// UpdateAuthorDTO defines the data transfer object for updating an existing author.
type UpdateAuthorDTO struct {
	Name string `json:"name" validate:"required"`
}

// AuthorRepository defines the methods that any
// data storage provider needs to implement to get
// and store authors
type AuthorRepository interface {
	FindAll(p *paginate.Param) ([]models.Author, int64, error)
	FindByID(id uint) (models.Author, error)
	Create(author models.Author) (models.Author, error)
	Update(author models.Author) (models.Author, error)
	Delete(id uint) error
}

// AuthorService defines the methods that an author service should implement
type AuthorService interface {
	GetAuthors(p *paginate.Param) ([]models.Author, int64, error)
	GetAuthor(id uint) (models.Author, error)
	CreateAuthor(dto CreateAuthorDTO) (models.Author, error)
	UpdateAuthor(id uint, dto UpdateAuthorDTO) (models.Author, error)
	DeleteAuthor(id uint) error
}

// authorService implements the AuthorService interface
type authorService struct {
	repo AuthorRepository
}

// NewAuthorService creates a new author service
func NewAuthorService(repo AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) GetAuthors(p *paginate.Param) ([]models.Author, int64, error) {
	authors, count, err := s.repo.FindAll(p)
	if err != nil {
		return nil, 0, err
	}
	return authors, count, nil
}

func (s *authorService) GetAuthor(id uint) (models.Author, error) {
	author, err := s.repo.FindByID(id)
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}

func (s *authorService) CreateAuthor(dto CreateAuthorDTO) (models.Author, error) {
	author := models.Author{
		Name: dto.Name,
	}
	createdAuthor, err := s.repo.Create(author)
	if err != nil {
		return models.Author{}, err
	}
	return createdAuthor, nil
}

func (s *authorService) UpdateAuthor(id uint, dto UpdateAuthorDTO) (models.Author, error) {
	// First, check if the author exists
	authorToUpdate, err := s.repo.FindByID(id)
	if err != nil {
		return models.Author{}, err // Author not found
	}

	authorToUpdate.Name = dto.Name

	updatedAuthor, err := s.repo.Update(authorToUpdate)
	if err != nil {
		return models.Author{}, err
	}
	return updatedAuthor, nil
}

func (s *authorService) DeleteAuthor(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
