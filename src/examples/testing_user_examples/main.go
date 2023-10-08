package main

import (
	"github.com/eduardor2m/go-mastery/src/examples/testing_user_examples/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	userHandler := handlers.UserHandler{}

	e.POST("/users", userHandler.Create)
	e.GET("/users", userHandler.FindAll)

	e.Logger.Fatal(e.Start(":8080"))
}
