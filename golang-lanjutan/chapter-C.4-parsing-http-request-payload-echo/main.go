package main 

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name string `json:"name" from:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()


	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	fmt.Println("Server started at localhost:9000")

	r.Start(":9000")
}