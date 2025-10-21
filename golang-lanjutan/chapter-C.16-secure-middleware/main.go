package main 

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/unrolled/secure"
)

func main() {

	secureMiddleware := secure.New(secure.options{
		AllowedHosts: []string{"localhost:9000", "www.google.com"},,
		FrameDeny: true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff: true,
		BrowserXssFilter: true,
	})

	e := echo.New()

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return c.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.Start(":9000", "server.crt", "server.key"))
}