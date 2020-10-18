package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/takenakasuji/go-layered-architecture/usecase"
)

type BookHandler interface {
	Index(c echo.Context) error
}

type bookHandler struct {
	bookUseCase usecase.BookUseCase
}

func NewBookHandler(bu usecase.BookUseCase) BookHandler {
	return &bookHandler{
		bookUseCase: bu,
	}
}

func (bh bookHandler) Index(c echo.Context) error {
	type request struct {
		Begin uint `query:"begin"`
		Limit uint `query:"limit"`
	}

	type bookField struct {
		Id       int64     `json:"id"`
		Title    string    `json:"title"`
		Author   string    `json:"author"`
		IssuedAt time.Time `json:"issued_at"`
	}

	type response struct {
		Books []bookField `json:"books"`
	}

	books, err := bh.bookUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	res := new(response)
	for _, book := range books {
		var bf bookField
		bf = bookField(*book)
		res.Books = append(res.Books, bf)
	}

	return c.JSON(http.StatusOK, res)
}
