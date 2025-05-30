package service_test

import (
	"testing"
	"log"
	"context"
	"github.com/lucasbadico/intensivo-first-service/repository"
	"github.com/lucasbadico/intensivo-first-service/service"
	"github.com/lucasbadico/intensivo-first-service/model"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb" // driver CouchDB
	"github.com/stretchr/testify/require"
	"github.com/tj/assert"
)

func build() *service.Service {
	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("erro ao criar cliente CouchDB: %s", err)
	}

	db := client.DB("notebook")
	if err := db.Err(); err != nil {
		log.Fatalf("erro ao conectar ao DB: %s", err)
	}

	repo := repository.New(db)
	srv := service.NewService(repo)
	return srv
}

func TestService(t *testing.T) {
	srv := build()
	if srv == nil {
		  t.Error("srv tem que existir")
	}

	t.Run("Service Create", func(t *testing.T) {
		// input\
		ctx := context.TODO()
		input := model.CreateNotebookInputDTO{
			Name: "Test Notebook",
			Description: "description notebook",
		}

		output, err := srv.Create(ctx, input)
		require.NoError(t, err)
		assert.Equal(t, input.Name, output.Name)
		assert.Equal(t, input.Description, output.Description)
		if output.ID == "" {
 				t.Error("ID tem que existir no output")
		}
	})
}