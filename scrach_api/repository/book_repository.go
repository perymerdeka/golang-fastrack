// storage of entity/model bean in system

package repository

import (
	"context"
	"scrach_api/models"
)


type BookRepository interface {
	Save(ctx context.Context, book models.Book) // save data
	Update(ctx context.Context, book models.Book) // update existing data
	Delete (ctx context.Context, bookId int) // delete data
	FindById(ctx context.Context, bookId int) (models.Book, error) // find single data by id
	FindAll(ctx context.Context) [] models.Book // find all data
}