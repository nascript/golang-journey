package services

import (
	entity "pustaka-api/entity"
	repositories "pustaka-api/repositories"
)

type Book = entity.Book
type AddBookRequest = entity.AddBookRequest
type UpdateBookRequest = entity.UpdateBookRequest
type Repository = repositories.Repository

type BookService interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book AddBookRequest) (Book, error)
	Update(ID int, book UpdateBookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
