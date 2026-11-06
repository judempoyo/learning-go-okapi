package services

import (
	"errors"

	"github.com/judempoyo/learning-go-okapi/go-modules-packages/models"
)

type BookRepository interface {
	FindAll() []models.Book
	FindByTitle(title string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	DeleteByTitle(title string) error
}

type BookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (bs *BookService) AddBookWithValidation(req *models.BookRequest) (*models.Book, error) {

	b, err := bs.repo.Create(models.Book{Title: req.Title, Price: req.Price})
	if err != nil {
		return nil, errors.New("Unable to create the book")
	}

	return &b, nil
}

func (bs *BookService) FindAllBooks() []models.Book {
	return bs.repo.FindAll()
}

func (bs *BookService) FindBook(title string) (*models.Book, error) {
	b, err := bs.repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}

	return &b, nil

}
func (bs *BookService) DeleteBookWithValidation(title string) error {

	err := bs.repo.DeleteByTitle(title)
	if err != nil {
		return err
	}

	return nil

}
