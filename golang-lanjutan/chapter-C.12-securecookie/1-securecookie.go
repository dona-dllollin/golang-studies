package main 

import (
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
	gubrak "github.com/novalagung/gubrak/v2"
	"net/http"
	"time"
	"fmt"
)

type M map[string]interface{}

var sc = securecookie.New([]byte("very-secret"), []byte("a-lot-secret-yay"))

func setCookie(c echo.Context, name string, data M) error {
	encoded, err := sc.Encode(name, data)

	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name: name,
		Value: encoded,
		Path: "/",
		Secure: false,
		HttpOnly: true,
		Expires: time.Now().Add(1 * time.Hour),
	}
	http.SetCookie(c.Response(), cookie)

	return nil
}

func getCookie(c echo.Context, name string) (M, error) {
	cookie, err := c.Request().Cookie(name)
	if err == nil {
		
		data := M{}

		if err = sc.Decode(name, cookie.Value, &data); err == nil {
			
			return data, nil
		}
	}

	return nil, err
}

func removeCookie(c echo.Context, name string) {
	cookie := &http.Cookie{}
	cookie.Name = name
	cookie.Path = "/"
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(c.Response(), cookie)
}

func main() {
	const CookieName = "data"
e := echo.New()
e.GET("/index", func(c echo.Context) error {
	data, err := getCookie(c, CookieName)
	if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
		return err
	}

	if data == nil {
		data = M{"Message": "Hello", "ID": gubrak.RandomString(32)}

		err = setCookie(c, CookieName, data)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, data)
})

e.GET("/delete", func(c echo.Context) error {
	removeCookie(c, CookieName)
	return c.Redirect(http.StatusTemporaryRedirect, "/index")
})

e.Logger.Fatal(e.Start(":9000"))

}