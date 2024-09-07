// to implementing business logic
package services

import (
	"context"
	"scrach_api/data/requests"
	"scrach_api/data/response"
)

type BookService interface {
	Create(ctx context.Context, request requests.BookCreateRequest)
	Update(ctx context.Context, request requests.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}
