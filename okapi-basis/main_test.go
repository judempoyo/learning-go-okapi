package main

import (
	"strings"
	"testing"

	"github.com/jkaninda/okapi"
	"github.com/jkaninda/okapi/okapitest"
)

func TestHealthHandler(t *testing.T) {
	server := okapi.NewTestServer(t)
	server.Get("/health", HealthHandler)

	okapitest.GET(t, server.BaseURL+"/health").
		ExpectStatusOK().
		ExpectBodyContains("Service healthy")
}

func TestGetBooksHandler(t *testing.T) {
	server := okapi.NewTestServer(t)
	server.Get("/books", GetBooksHandler)

	okapitest.GET(t, server.BaseURL+"/books").
		ExpectStatusOK().
		ExpectBodyContains("Go tour")
}

func TestGetBookHandler(t *testing.T) {
	server := okapi.NewTestServer(t)
	server.Get("/books/{id:int}", GetBookHandler)

	okapitest.GET(t, server.BaseURL+"/books/1").
		ExpectStatusOK().
		ExpectBodyContains("Go tour")

	okapitest.GET(t, server.BaseURL+"/books/999").
		ExpectStatusNotFound().
		ExpectBodyContains("Book not found")
}

func TestCreateBookHandler(t *testing.T) {
	server := okapi.NewTestServer(t)
	server.Post("/books", CreateBookHandler)

	req := BookRequest{
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
	server.Delete("/books/{id:int}", DeleteBookHandler)

	okapitest.DELETE(t, server.BaseURL+"/books/2").
		ExpectStatusOK().
		ExpectBodyContains("book deleted")
}

func TestCreateBookHandlerNegative(t *testing.T) {
	server := okapi.NewTestServer(t)
	server.Post("/books", CreateBookHandler)

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(map[string]any{"price": 50}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(BookRequest{Title: "Go", Price: 50}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(BookRequest{Title: "Advanced Go", Price: 0}).
		ExpectStatusBadRequest()

	okapitest.POST(t, server.BaseURL+"/books").
		JSONBody(BookRequest{Title: "Advanced Go", Price: -10}).
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
	server.Delete("/books/{id:int}", DeleteBookHandler)

	okapitest.DELETE(t, server.BaseURL+"/books/999").
		ExpectStatusNotFound().
		ExpectBodyContains("Book not found")
}