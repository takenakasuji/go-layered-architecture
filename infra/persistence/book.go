package persistence

import (
	"time"

	"github.com/takenakasuji/go-layered-architecture/domain/model"
	"github.com/takenakasuji/go-layered-architecture/domain/repository"
)

type bookPersistence struct{}

func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

func (bp bookPersistence) GetAll() ([]*model.Book, error) {
	book1 := model.Book{}
	book1.Id = 1
	book1.Title = "DDDが分かる本"
	book1.Author = "たろうくん"
	book1.IssuedAt = time.Now().Add(-24 * time.Hour)

	book2 := model.Book{}
	book2.Id = 2
	book2.Title = "レイヤードアーキテクチャが分かる本"
	book2.Author = "はなこさん"
	book2.IssuedAt = time.Now().Add(-24 * 7 * time.Hour)

	return []*model.Book{&book1, &book2}, nil
}
