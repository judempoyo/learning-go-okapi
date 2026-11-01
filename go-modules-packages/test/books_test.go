package test

import (
	"strings"
	"testing"

	"github.com/jkaninda/okapi"
	"github.com/jkaninda/okapi/okapitest"

	"github.com/judempoyo/learning-go-okapi/go-modules-packages/handlers"
	"github.com/judempoyo/learning-go-okapi/go-modules-packages/models"
	"github.com/judempoyo/learning-go-okapi/go-modules-packages/services"
	"github.com/judempoyo/learning-go-okapi/model"
)
func initiateStore () *services.BookService {

	bookStore := models.NewBookStore()
	bookService := services.NewBookService(bookStore)

	return  bookService
}

func TestGetBooksHandler(t *testing.T) {
	server := okapi.NewTestServer(t)
		bs := initiateStore()

	server.Get("/books", handlers.GetBooksHandler(bs))

	okapitest.GET(t, server.BaseURL+"/books").
		ExpectStatusOK().
		ExpectBodyContains("Go tour")
}

func TestGetBookHandler(t *testing.T) {
	server := okapi.NewTestServer(t)
		bs := initiateStore()

	server.Get("/books/{title:string}", handlers.GetBoookHandler(bs))

	okapitest.GET(t, server.BaseURL+"/books/Miabi").
		ExpectStatusOK().
		ExpectBodyContains("Miabi")

	okapitest.GET(t, server.BaseURL+"/books/test").
		ExpectStatusNotFound().
		ExpectBodyContains("Book not found")
}

func TestCreateBookHandler(t *testing.T) {
	server := okapi.NewTestServer(t)
		bs := initiateStore()

	server.Post("/books", handlers.CreateBookHandler(bs))

	req := model.BookRequest{
		Title: "Advanced Go Programming",
		Price: 50,
	}

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(req).
		ExpectStatusCreated().
		ExpectBodyContains("Advanced Go Programming")
}

func TestDeleteBookHandler(t *testing.T) {
	server := okapi.NewTestServer(t)

	bs := initiateStore()
	server.Delete("/books", handlers.DeleteBoookHandler(bs))

	req := model.BookRequest{
		Title: "Miabi",
	}

	okapitest.DELETE(t, server.BaseURL+"/books").
		JSONBody(req).
		ExpectStatusOK().
		ExpectBodyContains("Book succefully deleted")
}
func TestCreateBookHandlerNegative(t *testing.T) {
	server := okapi.NewTestServer(t)
	bs := initiateStore()

	server.Post("/books", handlers.CreateBookHandler(bs))

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(map[string]any{"price": 50}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(model.BookRequest{Title: "Go", Price: 50}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(model.BookRequest{Title: "Advanced Go", Price: 0}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(model.BookRequest{Title: "Advanced Go", Price: -10}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(map[string]any{}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		Body(strings.NewReader("{invalid-json}")).
		ExpectStatusBadRequest()
}

func TestDeleteBookHandlerNegative(t *testing.T) {
	server := okapi.NewTestServer(t)

	bs := initiateStore()
	server.Delete("/books", handlers.DeleteBoookHandler(bs))

	req := model.BookRequest{
		Title: "inexistantbook",
	}

	okapitest.DELETE(t, server.BaseURL+"/books").
		JSONBody(req).
		ExpectStatusNotFound().
		ExpectBodyContains("Book not found")
}