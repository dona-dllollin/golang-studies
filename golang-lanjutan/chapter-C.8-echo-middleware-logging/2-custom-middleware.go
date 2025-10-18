package main 

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func middlewareSomething(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("From custom middleware using net/http")
		next.ServeHTTP(w, r)
	})
}

func main() {
	e := echo.New()
	
	// middleware here
	e.Use(echo.WrapMiddleware(middlewareSomething))


	e.GET("/index", func(c echo.Context) (err error) {
		fmt.Println("Request received")

		return c.JSON(http.StatusOK, true)

	})

	e.Logger.Fatal(e.Start(":9000"))
}