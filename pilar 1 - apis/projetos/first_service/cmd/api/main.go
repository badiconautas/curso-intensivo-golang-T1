package main

import (
	"log"
	"net/http"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb" // driver CouchDB
	"github.com/gorilla/mux"
	"github.com/lucasbadico/intensivo-first-service/handler"
	"github.com/lucasbadico/intensivo-first-service/repository"
	"github.com/lucasbadico/intensivo-first-service/service"
)

func main() {
	log.Println("creating a webserver")

	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("erro ao criar cliente CouchDB: %s", err)
	}

	db := client.DB("notebooks")
	if err := db.Err(); err != nil {
		log.Fatalf("erro ao conectar ao DB: %s", err)
	}

	repo := repository.New(db)
	srv := service.NewService(repo)
	h := handler.New(srv)

	r := mux.NewRouter()

	h.MountHandlers(r)

	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", r)
}
