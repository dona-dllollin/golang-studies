package main 

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("From middleware One")
		return next(c)

	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("From middleware Two")
		return next(c)
	}
}

func main() {
	e := echo.New()
	
	// middleware here
	e.Use(middlewareOne)
	e.Use(middlewareTwo)


	e.GET("/index", func(c echo.Context) (err error) {
		fmt.Println("Request received")

		return c.JSON(http.StatusOK, true)

	})

	e.Logger.Fatal(e.Start(":9000"))
}