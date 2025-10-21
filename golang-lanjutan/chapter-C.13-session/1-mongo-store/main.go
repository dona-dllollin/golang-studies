// Deprecated Code: This code uses deprecated packages and is provided for reference only.
// package main

// import (
// 	"fmt"
// 	"github.com/gorilla/context"
// 	"github.com/kidstuff/mongostore"
// 	"github.com/labstack/echo"
// 	"gopkg.in/mgo.v2"
// 	"log"
// 	"net/http"
// 	"os"
// )

// const SESSION_ID = "id"

// func newMongoStore() *mongostore.MongoStore {
// 	mgoSession, err := mgo.Dial("localhost:27123")
// 	if err != nil {
// 		log.Println("ERROR", err)
// 		os.Exit(0)
// 	}

// 	dbCollection := mgoSession.DB("learnwebgolang").C("session")
// 	maxAge := 86400 * 7
// 	ensureTTL := true
// 	authKey := []byte("my-auth-key-very-secret")
// 	encryptionKey := []byte("my-encryption-key-very-secret123")

// 	store := mongostore.NewMongoStore(
// 		dbCollection,
// 		maxAge,
// 		ensureTTL,
// 		authKey,
// 		encryptionKey,
// 	)
// 	return store
// }

// func main() {
// 	store := newMongoStore()

// 	e := echo.New()

// 	e.Use(echo.WrapMiddleware(context.ClearHandler))

// 	e.GET("/set", func(c echo.Context) error {
// 		session, _ := store.Get(c.Request(), SESSION_ID)
// 		session.Values["message1"] = "hello"
// 		session.Values["message2"] = "world"
// 		session.Save(c.Request(), c.Response())

// 		return c.Redirect(http.StatusTemporaryRedirect, "/get")
// 	})

// 	e.GET("/get", func(c echo.Context) error {
// 		session, _ := store.Get(c.Request(), SESSION_ID)

// 		if len(session.Values) == 0 {
// 			return c.String(http.StatusOK, "empty result")
// 		}

// 		return c.String(http.StatusOK, fmt.Sprintf(
// 			"%s %s",
// 			session.Values["message1"],
// 			session.Values["message2"],
// 		))
// 	})

// 	e.GET("/delete", func(c echo.Context) error {
// 		session, _ := store.Get(c.Request(), SESSION_ID)
// 		session.Options.MaxAge = -1
// 		session.Save(c.Request(), c.Response())

// 		return c.Redirect(http.StatusTemporaryRedirect, "/get")
// 	})

// 	e.Logger.Fatal(e.Start(":9000"))
// }


package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const SESSION_ID = "session_id"

func newMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Println("❌ Gagal konek MongoDB:", err)
		os.Exit(1)
	}

	// Tes koneksi
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("❌ Tidak bisa ping MongoDB:", err)
		os.Exit(1)
	}

	log.Println("✅ Koneksi MongoDB berhasil!")
	return client
}

func main() {
	// Koneksi ke MongoDB
	client := newMongoClient()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("❌ Gagal disconnect:", err)
		}
	}()

	e := echo.New()

	// Middleware session
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret-key-very-secret"))))

	// Set session
	e.GET("/set", func(c echo.Context) error {
		sess, _ := session.Get(SESSION_ID, c)
		sess.Values["message1"] = "hello"
		sess.Values["message2"] = "world"
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	// Get session
	e.GET("/get", func(c echo.Context) error {
		sess, _ := session.Get(SESSION_ID, c)

		if len(sess.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}

		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			sess.Values["message1"],
			sess.Values["message2"],
		))
	})

	// Delete session
	e.GET("/delete", func(c echo.Context) error {
		sess, _ := session.Get(SESSION_ID, c)
		sess.Options.MaxAge = -1
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
