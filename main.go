package main

import (
    "github.com/jkaninda/okapi"
)

func main() {
    o := okapi.Default()
    
    o.Get("/", func(c *okapi.Context) error {
        return c.OK(okapi.M{"message": "Hello from Okapi!"})
    })
    
    if err := o.Start(); err != nil {
        panic(err)
    }
}