package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type customContext struct {
	echo.Context
}

func (c *customContext) HelloWorld() error {
	return c.String(http.StatusOK, "Hello, World!")
}

func helloWorld(c echo.Context) error {
	return c.(*customContext).HelloWorld()
}

func main() {
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&customContext{c})
		}
	})
	e.Get("/", helloWorld)

	e.Run(":1323")
}
