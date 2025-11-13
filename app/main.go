package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!")
	})

	e.GET("/goodbey", func(c echo.Context) error {
		return c.String(http.StatusOK, "Goodbey!")
	})

	e.Logger.Fatal(e.Start(":8081"))
}
