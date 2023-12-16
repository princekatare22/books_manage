package datastore

import (
	"books-manage/model"

	"gofr.dev/pkg/gofr"
)

type Book interface {
	GetByID(ctx *gofr.Context, id string) (*model.Book, error)
	Create(ctx *gofr.Context, model *model.Book) (*model.Book, error)
	Update(ctx *gofr.Context, model *model.Book) (*model.Book, error)
	Delete(ctx *gofr.Context, id int) error
}