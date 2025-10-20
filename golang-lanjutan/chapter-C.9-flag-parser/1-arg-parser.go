package main 

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	 "github.com/alecthomas/kingpin/v2"
)


var (
	argAppName = kingpin.Arg("name", "Application name").Required().String()
	argPort = kingpin.Arg("port", "Application port").Default("9000").Int()
)

func main() {
	kingpin.Parse()

	appName := *argAppName

	port := fmt.Sprintf(":%d", *argPort)

	fmt.Printf("Starting %s on port %s\n", appName, port)

	e := echo.New()

	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(port))
}