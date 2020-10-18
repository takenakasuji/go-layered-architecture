package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/takenakasuji/go-layered-architecture/handler/rest"
	"github.com/takenakasuji/go-layered-architecture/infra/persistence"
	"github.com/takenakasuji/go-layered-architecture/usecase"
)

func main() {
	bookPersistence := persistence.NewBookPersistence()
	bookUseCase := usecase.NewBookUseCase(bookPersistence)
	bookHandler := handler.NewBookHandler(bookUseCase)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/books", bookHandler.Index)

	e.Logger.Fatal(e.Start(":3000"))
}
