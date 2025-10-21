package main 

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	tmpl := template.Must(template.ParseGlob("*.html"))
	e := echo.New()
	const CSRFTokenHeader = "X-CSRF-Token"
	const CSRFKey = "csrf"

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRFTokenHeader,
		ContextKey: CSRFKey,
	}))

	e.GET("/index", func(c echo.Context) error {
 		data := make(M)
 		data[CSRFKey] = c.Get(CSRFKey)
		fmt.Println("CSRF token:", c.Get("csrf"))
 		return tmpl.Execute(c.Response(), data)
	})

	e.POST("/sayHello", func(c echo.Context) error {
 		data := make(M)
 		if err := c.Bind(&data); err != nil {
 		return err
 		}
 		message := fmt.Sprintf("hello %s", data["name"])
 		return c.JSON(http.StatusOK, message)
	})



	e.Logger.Fatal(e.Start(":8080"))
}