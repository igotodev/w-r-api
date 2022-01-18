package api

import "w-r-api/internal/domain/book/model"

type Service interface {
	GetByID(idStr string) *model.Book
	GetAll() *[]model.Book
	Create(book *model.Book) error
	Update(book *model.Book, idStr string) error
	Delete(idStr string) error
}
