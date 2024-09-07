package services

import (
	"context"
	"scrach_api/data/requests"
	"scrach_api/data/response"
	"scrach_api/helper"
	"scrach_api/models"
	"scrach_api/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

// Create implements BookService.
func (b *BookServiceImpl) Create(ctx context.Context, request requests.BookCreateRequest) {
	book := models.Book{
		Name: request.Name,
	}

	// save data to database with repository
	b.BookRepository.Save(ctx, book)
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)

	// proses delete book
	b.BookRepository.Delete(ctx, book.Id)
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)
	var booksResp []response.BookResponse
	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		booksResp = append(booksResp, book)
	}

	return booksResp
}

// FindById implements BookService.
func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)

	// return response for single book
	return response.BookResponse{Id: book.Id, Name: book.Name}
}

// Update implements BookService.
func (b *BookServiceImpl) Update(ctx context.Context, request requests.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	// proses update value
	book.Name = request.Name
	b.BookRepository.Update(ctx, book)

}
