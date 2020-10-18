package usecase

import (
	"github.com/takenakasuji/go-layered-architecture/domain/model"
	"github.com/takenakasuji/go-layered-architecture/domain/repository"
)

type BookUseCase interface {
	GetAll() ([]*model.Book, error)
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepository: br,
	}
}

func (bu bookUseCase) GetAll() (books []*model.Book, err error) {
	books, err = bu.bookRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}
