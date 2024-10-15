package bookservice

import (
	"go-api/store/book"
)

type BookService struct {
	BookStore bookstore.BookStore
}

func (s *BookService) ListBooks() ([]bookstore.Book, error) {
	return s.BookStore.GetAllBooks()
}