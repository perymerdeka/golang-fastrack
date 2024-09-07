//

package repository

import (
	"context"
	"database/sql"
	"errors"
	"scrach_api/helper"
	"scrach_api/models"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

// Delete implements BookRepository.
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// delete data from database
	SQL := "DELETE FROM books WHERE = $1"
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errExec)
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []models.Book {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// find all data from database
	SQL := "SELECT id, name FROM book" // query database
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	// iterate result
	var books []models.Book
	for result.Next() {
		book := models.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		// add to array
		books = append(books, book)
	}
	return books // return list of books
}

// FindById implements BookRepository.
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (models.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// query from database
	SQL := "SELECT id name FROM book WHERE id = $1"
	result, errQuery := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	// iterate result
	book := models.Book{}

	if result.Next() {
		// scan error pas proses iterasi
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		return book, nil // return single book
	} else {
		return book, errors.New("book ID was not found")
	}
}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(ctx context.Context, book models.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// query to database
	SQL := "INSERT INTO books(id, name) VALUES($1, $2)"
	_, errExec := tx.ExecContext(ctx, SQL, book.Id, book.Name)
	helper.PanicIfError(errExec)

}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(ctx context.Context, book models.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// query to database
	SQL := "UPDATE books SET name = $1 WHERE id = $2"
	_, errExec := tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(errExec)
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}