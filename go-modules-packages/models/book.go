package models

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type Book struct {
	ID    int `json:"id"`
	Title  string `json:"title" maxLength:"50" minLength:"5" required:"true" description:"Book name"`
	Price int    `json:"price" min:"1" max:"100" required:"true" description:"Book price"`
}

type BookRequest struct {
	Title string `json:"title" minLength:"3" example:"Learn Okapi framework" description:"Book title" required:"true"`
	Price int    `json:"price"  exclusiveMin:"0" description:"Book price" required:"true"`
}

type GetBookRequest struct {
	Title string `json:"title" required:"true"`
}

type Response struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
	Message any  `json:"details"`
	Data    any  `json:"data"`
}

type BookStore struct {
	books []Book
}

func NewBookStore() *BookStore {
	return &BookStore{
		books: []Book{
			{ID: 1, Title: "Go tour", Price: 100},
			{ID: 2, Title: "Miabi", Price: 50},
			{ID: 3, Title: "Apprendre posta", Price: 12},
		},
	}
}

func (bs *BookStore) FindAll() []Book {
	 books := make([]Book, len(bs.books))
    copy(books, bs.books)

    return books
}

func (bs *BookStore) FindByTitle(title string) (Book, error) {
	for _, b := range bs.books {
		if b.Title == title {
			return  b, nil
		}
	}

	return Book{}, errors.New("Book not found")
}
func (bs *BookStore) Create(b Book) (Book, error) {

	_, err := bs.FindByTitle(b.Title)
	if err == nil {
		return Book{}, errors.New("Similar book already exist")
	}
	id, _ := rand.Int(rand.Reader, big.NewInt(100))
	b.ID = int(id.Int64())
	bs.books = append(bs.books, b)
	return b,nil
}
func (bs *BookStore) DeleteByTitle(title string) error {
	for i, b := range bs.books {
		if b.Title == title {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)
			return nil
		}
	}

	return errors.New("Book not found")

}
/* 
func (s *BookStore)	Add (b Book) {
	s.books = append(s.books, b)
}

func (s *BookStore) Delete(index int) {
  
	s.books = append(s.books[:index], s.books[index+1:]...)
}

func (s *BookStore) Search(title string) (int){
	index := -1
	for i, b := range s.books {
		if b.Title == title {
			index = i
		}
	}

	return index

} */