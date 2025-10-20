package main 

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"net/http"
)


func main() {
	e := echo.New()
 
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	if err := viper.ReadInConfig(); err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Print("Starting ", viper.GetString("appName"))
	e.Logger.Print(e.Start(":" + viper.GetString("server.port")))
}