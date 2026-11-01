package models

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
	Books []Book
}

func NewBookStore() *BookStore {
	return &BookStore{
		Books: []Book{
			{ID: 1, Title: "Go tour", Price: 100},
			{ID: 2, Title: "Miabi", Price: 50},
			{ID: 3, Title: "Apprendre posta", Price: 12},
		},
	}
}

func (s *BookStore)	Add (b Book) {
	s.Books = append(s.Books, b)
}

func (s *BookStore) Delete(index int) {
  
	s.Books = append(s.Books[:index], s.Books[index+1:]...)
}

func (s *BookStore) Search(title string) (int){
	index := -1
	for i, b := range s.Books {
		if b.Title == title {
			index = i
		}
	}

	return index

}