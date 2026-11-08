package handlers

import (
	"net/http"

	"github.com/jkaninda/okapi"

	"github.com/judempoyo/learning-go-okapi/go-modules-packages/models"
	"github.com/judempoyo/learning-go-okapi/go-modules-packages/services"
)

func GetBooksHandler(bs *services.BookService) func(*okapi.Context) error {
	return func(c *okapi.Context) error {
		return c.OK(models.Response{
			Success: true,
			Status:  http.StatusOK,
			Message: "All books retrieved succefully",
			Data:    bs.FindAllBooks(),
		})
	}
}
func GetBoookHandler(bs *services.BookService) func(*okapi.Context) error {
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

		b, err := bs.FindBook(bookTitle)

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
func DeleteBoookHandler(bs *services.BookService) func(*okapi.Context) error {
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
		err := bs.DeleteBookWithValidation(req.Title)

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
			Message: "Book succefully deleted",
			Data:    nil,
		}
		return c.OK(response)
	}
}
func CreateBookHandler(bs *services.BookService) func(*okapi.Context) error {
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

		book, err := bs.AddBookWithValidation(&req)
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
