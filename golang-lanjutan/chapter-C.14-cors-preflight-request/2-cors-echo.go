package main 

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string("https://www.google.com", "https://novalagung.com"),
		AllowedMethods: []string{"OPTIONS", "POST", "GET", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,

	})

	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/index", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.Logger.Fatal(e.Start(":9000"))
}