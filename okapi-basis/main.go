package main

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/jkaninda/okapi"

	"github.com/judempoyo/learning-go-okapi/model"
)

func generateID(req *model.BookRequest) model.Book {
	id, _ := rand.Int(rand.Reader, big.NewInt(100))

	return model.Book{
		ID:    id.String(),
		Name:  req.Title,
		Price: req.Price,
	}
}

func CreateBookHandler(c *okapi.Context) error {
	var req model.BookRequest
	if err := c.Bind(&req); err != nil {
		return c.ErrorBadRequest(model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	book := generateID(&req)
	books = append(books, book)

	response := model.Response{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Book created succefully",
		Data:    book,
	}
	return c.Created(response)

}
func HealthHandler(c *okapi.Context) error {
	return c.OK(model.Response{
		Success: true,
		Status:  http.StatusOK,
		Message: "Service healthy",
		Data:    nil,
	})
}

func GetBooksHandler(c *okapi.Context) error {

	return c.OK(model.Response{
		Success: true,
		Status:  http.StatusOK,
		Message: "All books retrieved succefully",
		Data:    books,
	})
}

func GetBookHandler(c *okapi.Context) error {
	bookId := c.Param("id")

	if bookId == "" {
		return c.ErrorBadRequest(model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Missing book id",
			Data:    nil,
		})
	}

	index := -1

	for i, b := range books {
		if b.ID == bookId {
			index = i
			break
		}
	}

	if index == -1 {
		return c.ErrorNotFound(model.ErrorResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Book not found",
			Data:    nil,
		})

	}
	book := books[index]

	return c.OK(model.Response{
		Success: true,
		Status:  http.StatusOK,
		Message: "Book retrieved succefully",
		Data:    book,
	})
}

func DeleteBookHandler(c *okapi.Context) error {
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

	if indexToDelete == -1 {
		return c.ErrorNotFound(model.ErrorResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Book not found",
			Data:    nil,
		})

	}

	books = append(books[:indexToDelete], books[indexToDelete+1:]...)

	return c.OK(okapi.M{
		"Success": true,
		"Message": "book deleted",
		"Data":    books,
	})
}

var books = []model.Book{
	{ID: "1", Name: "Go tour", Price: 100},
	{ID: "2", Name: "Miabi ", Price: 50},
	{ID: "3", Name: "Apprendre posta", Price: 12},
}

func main() {
	o := okapi.Default()

	o.Post("/books", CreateBookHandler, okapi.DocSummary("Create a new Book"),
		okapi.DocRequestBody(model.BookRequest{}),
		okapi.DocResponse(model.Response{}),
		okapi.DocResponse(http.StatusBadRequest, model.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, model.ErrorResponse{}),
	)

	o.Delete("/books/{id:int}", DeleteBookHandler, okapi.DocResponse(model.Response{}),
		okapi.DocResponse(http.StatusBadRequest, model.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, model.ErrorResponse{}))
	o.Get("/health", HealthHandler,

		okapi.DocResponse(model.Response{}),
		okapi.DocResponse(http.StatusBadRequest, model.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, model.ErrorResponse{}))

	o.Get("books", GetBooksHandler,
		okapi.DocResponse(model.Response{}),
		okapi.DocResponse(http.StatusBadRequest, model.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, model.ErrorResponse{}),
	)

	o.Get("/books/{id:int}", GetBookHandler,

		okapi.DocResponse(model.Response{}),
		okapi.DocResponse(http.StatusBadRequest, model.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, model.ErrorResponse{}))

	if err := o.Start(); err != nil {
		panic(err)
	}

}
