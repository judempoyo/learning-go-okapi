package handlers

import (
	"net/http"

	"github.com/jkaninda/okapi"

	"github.com/judempoyo/learning-go-okapi/go-modules-packages/models"
	"github.com/judempoyo/learning-go-okapi/go-modules-packages/services"
)

func GetBooksHandler(bookService *services.BookService) func(*okapi.Context) error {
	return func(c *okapi.Context) error {
		return c.OK(models.Response{
			Success: true,
			Status:  http.StatusOK,
			Message: "All books retrieved succefully",
			Data:    bookService.Store.Books,
		})
	}
}
func GetBoookHandler(bookservices *services.BookService) func(*okapi.Context) error {
	return func(c *okapi.Context) error {

		bookTitle := c.Param("title")

		if bookTitle == "" {
			return c.ErrorBadRequest(models.ErrorResponse{
				Success: false,
				Status:  http.StatusBadRequest,
				Message: "Missing book title",
				Data:    nil,
			})
		}
		
		b, err := bookservices.FindBook(bookTitle)

		if err != nil {
			return c.ErrorNotFound(models.ErrorResponse{
				Success: false,
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
		}

		response := models.Response{
			Success: true,
			Status:  http.StatusOK,
			Message: "Book retrieved succefully",
			Data:    b,
		}
		return c.OK(response)
	}
}
func DeleteBoookHandler(bookservices *services.BookService) func(*okapi.Context) error {
	return func(c *okapi.Context) error {
		var req models.GetBookRequest
		if err := c.Bind(&req); err != nil {
			return c.ErrorBadRequest(models.ErrorResponse{
				Success: false,
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
		}
		msg, err := bookservices.DeleteBookWithValidation(req.Title)

		if err != nil {
			return c.ErrorNotFound(models.ErrorResponse{
				Success: false,
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
		}

		response := models.Response{
			Success: true,
			Status:  http.StatusOK,
			Message: msg,
			Data:    nil,
		}
		return c.OK(response)
	}
}
func CreateBookHandler(bookService *services.BookService) func(*okapi.Context) error {
	return func(c *okapi.Context) error {
		var req models.BookRequest
		if err := c.Bind(&req); err != nil {
			return c.ErrorBadRequest(models.ErrorResponse{
				Success: false,
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
		}

		book, err := bookService.AddBookWithValidation(&req)
		if err != nil {
			return c.ErrorBadRequest(models.ErrorResponse{
				Success: false,
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
		}

		response := models.Response{
			Success: true,
			Status:  http.StatusCreated,
			Message: "Book created succefully",
			Data:    book,
		}
		return c.Created(response)
	}
}
