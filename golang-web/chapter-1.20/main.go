package main 

import (
	"chapter-1.20/conf"
	"fmt"
	"net/http"
	"time"
	"log"
)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Shared().Log.Verbose {
		log.Println("Incoming request from", r.Host, "accessing", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	router := new(CustomMux)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	 router.HandleFunc("/howareyou", func(w http.ResponseWriter, r *http.Request) {
 		w.Write([]byte("How are you?"))
 	})


	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = conf.Shared().Server.ReadTimeout * time.Second
	server.WriteTimeout = conf.Shared().Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%d", conf.Shared().Server.Port)

	if conf.Shared().Log.Verbose {
		log.Printf("Server started at %s\n", server.Addr)
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}