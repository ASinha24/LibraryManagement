package bookStore

import (
	"context"
	"errors"
	"sync"

	"github.com/ASinha24/LibraryManagementSystem/bookStore/model"
)

type BookStore interface {
	AddBook(ctx context.Context, book *model.Book) (*model.Book, error)
	UpdateBook(ctx context.Context, book *model.Book) (*model.Book, error)
	DeleteBook(ctx context.Context, id string) error
	FindBookByID(ctx context.Context, id string) (*model.Book, error)
	FindBookByName(ctx context.Context, name string) (*model.Book, error)
	GetAllBooks(ctx context.Context) ([]*model.Book, error)
	Close()
}

type BookStoreInMem struct {
	books  map[string]*model.Book
	bookch chan<- *model.Book
	once   sync.Once
}

func NewBookStore() BookStore {
	bookch := make(chan *model.Book)
	s := &BookStoreInMem{books: make(map[string]*model.Book), bookch: bookch}
	go func(ch <-chan *model.Book) {
		for i := range ch {
			s.books[i.ID] = i
		}
	}(bookch)
	return s
}

func (b BookStoreInMem) Close() {
	b.once.Do(func() { close(b.bookch) })
}

func (b BookStoreInMem) AddBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	b.bookch <- book
	return book, nil
}

func (b BookStoreInMem) UpdateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	oldBook, ok := b.books[book.ID]
	if !ok {
		return nil, errors.New("Book not found in the library")
	}
	oldBook.Name = book.Name
	oldBook.Quantity = book.Quantity
	b.bookch <- oldBook
	return book, nil
}

func (b BookStoreInMem) DeleteBook(ctx context.Context, id string) error {
	_, ok := b.books[id]
	if !ok {
		return errors.New("Book not fount in the Library")
	}
	delete(b.books, id)
	return nil
}

func (b BookStoreInMem) FindBookByID(ctx context.Context, id string) (*model.Book, error) {
	book, ok := b.books[id]
	if !ok {
		return nil, errors.New("Book not fount in the Library")
	}
	return book, nil
}

func (b BookStoreInMem) FindBookByName(ctx context.Context, name string) (*model.Book, error) {
	for _, m := range b.books {
		if m.Name == name {
			return m, nil
		}
	}
	return nil, errors.New("Book not fount in the Library")
}

func (b BookStoreInMem) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	for _, b := range b.books {
		books = append(books, b)
	}
	return books, nil
}
