package services

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/judempoyo/learning-go-okapi/go-modules-packages/models"
)

type BookService struct {
	Store *models.BookStore
}

func NewBookService(store *models.BookStore) *BookService {
	return &BookService{Store: store}
}

func GenerateID(req *models.BookRequest) models.Book {
	id, _ := rand.Int(rand.Reader, big.NewInt(100))

	return models.Book{
		ID:    int(id.Int64()),
		Title: req.Title,
		Price: req.Price,
	}
}

func (s *BookService) AddBookWithValidation(req *models.BookRequest) (*models.Book, error) {
	for _, b := range s.Store.Books {
		if b.Title == req.Title {
			return nil, errors.New("this book exist")
		}
	}

	b := GenerateID(req)

	s.Store.Add(b)

	println(s)
	return &b, nil
}

func (s *BookService) FindBook(title string) (*models.Book, error) {
	i := s.Store.Search(title)

	if i == -1 {	
		return nil , errors.New("Book not found")
	}

	b := s.Store.Books[i]

	return &b, nil

}
func (s *BookService) DeleteBookWithValidation(title string) (string, error) {

	i := s.Store.Search(title)

	if i == -1 {	
		return "" , errors.New("Book not found")
	}

	s.Store.Delete(i)

	return "Book succefully deleted", nil

	
}
