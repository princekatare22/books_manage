package datastore

import (
	"database/sql"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"books-manage/model"
)

type book struct{}

func New() *book {
	return &book{}
}

func (s *book) GetByID(ctx *gofr.Context, id string) (*model.Book, error) {
	var resp model.Book

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,name,price,genre FROM books where id=?", id).
		Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Genre)
	switch err {
	case sql.ErrNoRows:
		return &model.Book{}, errors.EntityNotFound{Entity: "book", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Book{}, err
	}
}

func (s *book) Update(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE books SET name=?,price=?,genre=? WHERE id=?",
		book.Name, book.Price, book.Genre, book.ID)
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return book, nil
}

func (s *book) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM books where id=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (s *book) Create(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	var resp model.Book

	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO books (name, price, genre) VALUES ( ?,?,?)", book.Name, book.Price, book.Genre)
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}
	err = ctx.DB().QueryRowContext(ctx, "SELECT id, name, price, genre FROM books WHERE id = ?", lastInsertID).Scan(&resp.ID, &resp.Name, &resp.Price, &resp.Genre)
	if err != nil {
		return &model.Book{}, errors.DB{Err: err}
	}

	return &resp, nil
}