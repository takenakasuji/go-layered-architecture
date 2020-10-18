package repository

import (
	"github.com/takenakasuji/go-layered-architecture/domain/model"
)

type BookRepository interface {
	GetAll() ([]*model.Book, error)
}
