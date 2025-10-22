package main 

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"io"
)

func main() {
	e := echo.New()

	e.Use(middleware.Gzip())

	e.GET("/image", func(c echo.Context) error {
		f, err := os.Open("sample.png")
		if f == nil {
			defer f.Close()
		}
		if err != nil {
			return err
		}

		_, err = io.Copy(c.Response(), f)
		if err != nil {
			return err
		}
		return nil

	})

	e.Logger.Fatal(e.Start(":9000"))

}