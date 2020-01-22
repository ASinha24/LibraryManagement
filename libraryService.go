package LibraryManagementSystem

import (
	"context"

	"github.com/ASinha24/LibraryManagementSystem/api"
	"github.com/ASinha24/LibraryManagementSystem/bookStore"
	"github.com/ASinha24/LibraryManagementSystem/bookStore/model"
	"github.com/pborman/uuid"
)

type LibraryManager interface {
	AddBook(ctx context.Context, book *api.BookRequest) (*api.BookCreateResponse, error)
	UpdateBook(ctx context.Context, bookId string, book *api.BookRequest) (*api.BookCreateResponse, error)
	DeleteBook(ctx context.Context, id string) error
}

type libService struct {
	bookstore bookStore.BookStore
}

func NewBookService(bookStore bookStore.BookStore) *libService {
	return &libService{
		bookstore: bookStore,
	}
}

func (l *libService) AddBook(ctx context.Context, book *api.BookRequest) (*api.BookCreateResponse, error) {
	newBook, err := l.bookstore.AddBook(ctx, &model.Book{
		ID:       uuid.NewUUID().String(),
		Name:     book.Name,
		Quantity: book.Quantity,
	})
	if err != nil {
		return nil, &api.BookError{Code: api.BookCreationFailed, Message: "cannot add new book", Description: err.Error()}
	}
	return &api.BookCreateResponse{ID: newBook.ID, BookRequest: book}, nil
}

func (l *libService) UpdateBook(ctx context.Context, bookId string, book *api.BookRequest) (*api.BookCreateResponse, error) {
	updateBook, err := l.bookstore.UpdateBook(ctx, &model.Book{
		ID:       bookId,
		Name:     book.Name,
		Quantity: book.Quantity,
	})
	if err != nil {
		return nil, &api.BookError{Code: api.BookNotFound, Message: "not able to update book", Description: err.Error()}
	}
	return &api.BookCreateResponse{ID: updateBook.ID, BookRequest: book}, nil
}

func (l *libService) DeleteBook(ctx context.Context, id string) error {
	err := l.bookstore.DeleteBook(ctx, id)
	if err != nil {
		return &api.BookError{Code: api.BookDeletionFailed, Message: "id not exist", Description: err.Error()}
	}
	return nil
}
