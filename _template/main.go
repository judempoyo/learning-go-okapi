package main

import (
	"net/http"

	"github.com/jkaninda/okapi"
)

func main() {
	o := okapi.Default()

	o.Get("/", func(c *okapi.Context) error {
		return c.JSON(http.StatusOK, okapi.M{"message": "Hello Okapi"})
	})

	if err := o.Start(); err != nil {
		panic(err)
	}
}