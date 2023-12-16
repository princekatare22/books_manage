package main

import (
	"books-manage/datastore"
	"books-manage/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/books/{id}", h.GetByID)
	app.POST("/books", h.Create)
	app.PUT("/books/{id}", h.Update)
	app.DELETE("/books/{id}", h.Delete)

	app.Server.HTTP.Port = 9092
	app.Start()
}
