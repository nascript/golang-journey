package services

import (
	entity "pustaka-api/entity"
	repositories "pustaka-api/repositories"
)

type Book = entity.Book
type BookRequest = entity.BookRequest
type Repository = repositories.Repository

type BookService interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
