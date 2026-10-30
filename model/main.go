package model

type Book struct {
	ID    string `json:"id"`
	Name  string `json:"title" maxLength:"50" minLength:"5" required:"true" description:"Book name"`
	Price int    `json:"price" min:"1" max:"100" required:"true" description:"Book price"`
}

type BookRequest struct {
	Title string `json:"title" minLength:"3" example:"Learn Okapi framework" description:"Book title" required:"true"`
	Price int    `json:"price"  exclusiveMin:"0" description:"Book price" required:"true"`
}

type GetBookRequest struct {
	ID int `json:"id" required:"true"`
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