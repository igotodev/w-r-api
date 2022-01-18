package service

import (
	"w-r-api/internal/controllers/api"
	"w-r-api/internal/domain/book/model"
	"w-r-api/internal/domain/book/storage"
)

type service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) api.Service {
	return &service{storage: storage}
}

func (s *service) GetByID(idStr string) *model.Book {
	return s.storage.GetByID(idStr)
}

func (s *service) GetAll() *[]model.Book {
	return s.storage.GetAll()
}

func (s *service) Create(book *model.Book) error {
	return s.storage.Create(book)
}

func (s *service) Update(book *model.Book, idStr string) error {
	return s.storage.Update(book, idStr)
}

func (s *service) Delete(idStr string) error {
	return s.storage.Delete(idStr)
}
