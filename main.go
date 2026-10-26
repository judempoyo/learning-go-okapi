package main

import (
	"crypto/rand"
	"math/big"

	"github.com/jkaninda/okapi"
)

type Book struct {
	ID    string  	`json:"id"`
	Title string  	`json:"title"`
	Price int 		`json:"price"`
}

type BookRequest struct {
	Title string  	`json:"title" minLength:"3" example:"Learn Okapi framework" description:"Book title"`
	Price int 		`json:"price"  exclusiveMin:"0" description:"Book price"`
}

func (req *BookRequest) GenerateId() Book {
	id, _ := rand.Int(rand.Reader, big.NewInt(100))

	

	return Book{
		id.String(),
		req.Title,
		req.Price,
	}
}

func main() {
	o := okapi.Default()
	books := []Book{
		{"1", "Go tour", 100},
		{"2", "Miabi ", 50},
		{"3", "Apprendre posta", 12},
	}

	o.Post("/books", func(c *okapi.Context) error {
		var req BookRequest
		if err := c.Bind(&req); err != nil {
			return c.AbortBadRequest("Invalid request", err)
		}

		book := req.GenerateId()
		books = append(books, book)

		return c.Created(book)


	}, okapi.DocSummary("Create a new Book"),
		okapi.DocRequestBody(BookRequest{}),
	)

	o.Delete("/books/{id:string}", func(c *okapi.Context) error {
		bookId := c.Param("id")

		if bookId == "" {
			return c.ErrorBadRequest("missing book id")
		}

		indexToDelete := -1

		for i, b := range books {
			if b.ID == bookId {
				indexToDelete = i
				break
			}
		}

		if indexToDelete != -1 {
			books = append(books[:indexToDelete], books[indexToDelete+1:]...)

		}
		return c.OK(okapi.M{
			"Success": true,
			"Message": "book deleted",
			"Data":    books,
		})
	})
	o.Get("/health", func(c *okapi.Context) error {
		return c.OK(okapi.M{
			"message": "Hello Its  work!",
		})
	})

	o.Get("books", func(c *okapi.Context) error {

		return c.OK(books)
	})

	if err := o.Start(); err != nil {
		panic(err)
	}

}
