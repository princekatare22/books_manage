package datastore

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"books-manage/model"
)

type book struct{}

func New() *book {
	return &book{}
}

func (s *book) GetByID(ctx *gofr.Context, id int) (*model.Book, error) {
	var resp model.Book

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,name,price,genre FROM books where id=$1", id).
		Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Genre)
	switch err {
	case sql.ErrNoRows:
		return &model.Book{}, errors.EntityNotFound{Entity: "book", ID: strconv.Itoa(id)}
	case nil:
		return &resp, nil
	default:
		return &model.Book{}, err
	}
}

func (s *book) Create(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	var resp model.Book

	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO books (name, price, genre) VALUES($1,$2,$3)"+
		" RETURNING  id,name,price,genre", book.ID, book.Name, book.Price, book.Genre).Scan(
		&resp.ID, &resp.Name, &resp.Price, &resp.Genre)

	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *book) Update(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE books SET name=$1,price=$2,genre=$3 WHERE id=$4",
		book.Name, book.Price, book.Genre, book.ID)
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return book, nil
}

func (s *book) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM books where id=$1", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}