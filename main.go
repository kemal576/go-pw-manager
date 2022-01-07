package main

import (
	"log"
	"net/http"

	"github.com/kemal576/go-pw-manager/internal/router"
	"github.com/kemal576/go-pw-manager/repository"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	dbconn, err := repository.Conn()
	if err != nil {
		log.Panic(err)
	}
	defer dbconn.Close()

	db := repository.New(dbconn)
	router := router.New(*db)
	handler := cors.AllowAll().Handler(router.Router)

	println("Server started: http://localhost:5764")
	log.Fatal(http.ListenAndServe(":5764", handler))
}
