package main

import (
	"net/http"

	"github.com/jkaninda/okapi"

	"github.com/judempoyo/learning-go-okapi/go-modules-packages/handlers"
	"github.com/judempoyo/learning-go-okapi/go-modules-packages/models"
	"github.com/judempoyo/learning-go-okapi/go-modules-packages/services"
)

func main() {
	o := okapi.Default()

	o.Get("/", func(c *okapi.Context) error {
		return c.JSON(http.StatusOK, okapi.M{"message": "Hello Okapi"})
	})

	bookStore := models.NewBookStore()
	bookService := services.NewBookService(bookStore)

	o.Post("/books", handlers.CreateBookHandler(bookService), okapi.DocSummary("Create a new Book"),
		okapi.DocRequestBody(models.BookRequest{}),
		okapi.DocResponse(models.Response{}),
		okapi.DocResponse(http.StatusBadRequest, models.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, models.ErrorResponse{}),
	)
	o.Delete("/books", handlers.DeleteBoookHandler(bookService), okapi.DocSummary("Delete a Book"),
		okapi.DocRequestBody(models.GetBookRequest{}),
		okapi.DocResponse(models.Response{}),
		okapi.DocResponse(http.StatusBadRequest, models.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, models.ErrorResponse{}),
	)

	o.Get("/books/{title: string}", handlers.GetBoookHandler(bookService), okapi.DocSummary("Get a Book"),
		okapi.DocResponse(models.Response{}),
		okapi.DocResponse(http.StatusBadRequest, models.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, models.ErrorResponse{}),
	)
	o.Get("/books", handlers.GetBooksHandler(bookService),
		okapi.DocResponse(models.Response{}),
		okapi.DocResponse(http.StatusBadRequest, models.ErrorResponse{}),
		okapi.DocResponse(http.StatusInternalServerError, models.ErrorResponse{}),
	)
	if err := o.Start(); err != nil {
		panic(err)
	}
}
