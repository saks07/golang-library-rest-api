package borrowedservice

import (
	"go-api/store/borrowed_books"
)

type BorrowedService struct {
	BorrowedStore borrowedstore.BorrowedStore
}

func (s *BorrowedService) ListBorrowedBooks(userId string) ([]borrowedstore.BorrowedBook, error) {
	return s.BorrowedStore.GetUserBorrowedBooks(userId)
}

func (s *BorrowedService) ListReturnedBooks(userId string) ([]borrowedstore.BorrowedBook, error) {
	return s.BorrowedStore.GetUserReturnedBooks(userId)
}

func (s *BorrowedService) CreateBorrowedBooks(bookId int, userId int, borrowDate string) error {
	return s.BorrowedStore.SaveBorrowedBooks(bookId, userId, borrowDate)
}

func (s *BorrowedService) UpdateBorrowedBooks(id int, returnDate string) error {
	return s.BorrowedStore.UpdateReturnedBooks(id, returnDate)
}