package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var store *sessions.CookieStore

func main() {
	e := echo.New()

	// --- connect PostgreSQL ---
	db, err := sql.Open("pgx", "postgres://user:password@localhost:5432/learnwebgolang?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// --- session (pakai cookie) ---
	store = sessions.NewCookieStore([]byte("super-secret-key"))
	e.Use(session.Middleware(store))

	e.GET("/set", func(c echo.Context) error {
		sess, _ := session.Get("session_id", c)
		sess.Values["message1"] = "hello"
		sess.Values["message2"] = "world"
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.GET("/get", func(c echo.Context) error {
		sess, _ := session.Get("session_id", c)
		if len(sess.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}
		return c.String(http.StatusOK, fmt.Sprintf("%s %s", sess.Values["message1"], sess.Values["message2"]))
	})

	e.GET("/delete", func(c echo.Context) error {
		sess, _ := session.Get("session_id", c)
		sess.Options.MaxAge = -1
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
