package main

import (
	"log"
	"net/http"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb" // driver CouchDB
	"github.com/gorilla/mux"
	"github.com/lucasbadico/intensivo-hello-world/handler"
)

func main() {
	log.Println("creating a webserver")

	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("erro ao criar cliente CouchDB: %s", err)
	}

	db := client.DB("notebook")
	if err := db.Err(); err != nil {
		log.Fatalf("erro ao conectar ao DB: %s", err)
	}

	r := mux.NewRouter()

	h := handler.New(db)

	r.HandleFunc("/health", h.Health).Methods("GET")

	r.HandleFunc("/notebooks", h.Create).Methods("POST")

	r.HandleFunc("/notebooks/{notebook_id}", h.Get).Methods("GET")
	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", r)
}
