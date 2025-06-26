package service_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/lucasbadico/surfbook-api-v1/foundation"
	"github.com/lucasbadico/surfbook-api-v1/model"
	"github.com/lucasbadico/surfbook-api-v1/repository"
	"github.com/lucasbadico/surfbook-api-v1/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"
)

func createUser(db *sql.DB, id uuid.UUID, name, email string) error {
	query := `INSERT INTO users (user_id, name, email) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, id, name, email)
	return err
}

func setup(conn string) (*service.NotebookService, uuid.UUID) {
	userID := uuid.New()
	db, err := foundation.NewPostgresDB(conn)
	if err != nil {
		log.Fatalf("Conexão com PostgreSQL não executada! %v", err)
	}
	fmt.Println("Conexão com PostgreSQL estabelecida com sucesso.")

	// need a user in db to complain with foreing key
	createUser(db, userID, "test user", "test@test.com")

	repo := repository.NewNotebookRepository(db)
	srv := service.NewNotebookService(repo)

	return srv, userID
}

func TestNotebookService(t *testing.T) {
	srv, userID := setup("postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable")
	if srv == nil {
		t.Error("srv tem que existir")
	}

	t.Run("Create Notebook", func(t *testing.T) {
		// input\
		ctx := context.TODO()
		input := model.CreateNotebookDTO{
			UserID:      userID,
			Icon:        "folder",
			Name:        "Test Notebook",
			Image:       "google.com/image",
			Description: "description notebook",
		}

		notebook, err := srv.Create(ctx, input)
		require.NoError(t, err)

		t.Run("it should an id", func(t *testing.T) {
			require.NotNil(t, notebook.NotebookID)
			// assert.Equal(t, input.Name, output.Name)
			// assert.Equal(t, input.Description, output.Description)
			// if output.ID == "" {
			// 	t.Error("ID tem que existir no output")
			// }

		})

		t.Run("it should create at and update at", func(t *testing.T) {
			require.NotEmpty(t, notebook.CreatedAt)
			require.NotEmpty(t, notebook.UpdatedAt)

		})

		t.Run("it should not have deleted at", func(t *testing.T) {
			require.Empty(t, notebook.DeletedAt)

		})

	})

	t.Run("List Notebook", func(t *testing.T) {
		ctx := context.TODO()
		input := model.CreateNotebookDTO{
			UserID:      userID,
			Icon:        "folder",
			Name:        "Test Notebook",
			Image:       "google.com/image",
			Description: "description notebook",
		}

		_, err := srv.Create(ctx, input)
		require.NoError(t, err)

		notebooks, err := srv.ListFromUser(ctx, model.ListNotebookFromUserDTO{
			UserID: userID,
		})
		require.NoError(t, err)

		t.Run("it should 2", func(t *testing.T) {
			require.Len(t, notebooks, 2)
		})
	})

	t.Run("Find Notebook", func(t *testing.T) {
		// input\
		ctx := context.TODO()
		input := model.CreateNotebookDTO{
			UserID:      userID,
			Icon:        "folder",
			Name:        "Test Notebook",
			Image:       "google.com/image",
			Description: "description notebook",
		}

		created, err := srv.Create(ctx, input)
		require.NoError(t, err)

		found, err := srv.FindNotebookFromUser(ctx, model.FindNotebookFromUserDTO{
			UserID:     userID,
			NotebookID: created.NotebookID,
		})
		require.NoError(t, err)

		t.Run("it have same ids", func(t *testing.T) {
			assert.Equal(t, created.NotebookID, found.NotebookID)
		})

		t.Run("it have same description and icon", func(t *testing.T) {
			assert.Equal(t, created.Description, found.Description)
			assert.Equal(t, created.Icon, found.Icon)
			assert.Equal(t, created.Name, found.Name)
			assert.Equal(t, created.Image, found.Image)
		})

		// needs parse on date
		// t.Run("it have same created at and diff updated at", func(t *testing.T) {
		// 	assert.Equal(t, created.CreatedAt, updated.CreatedAt)
		// 	assert.NotEqual(t, created.UpdatedAt, updated.UpdatedAt)
		// })

	})

	t.Run("Delete Notebook", func(t *testing.T) {
		ctx := context.TODO()
		input := model.CreateNotebookDTO{
			UserID:      userID,
			Icon:        "folder",
			Name:        "Test Notebook",
			Image:       "google.com/image",
			Description: "description notebook",
		}

		created, err := srv.Create(ctx, input)
		require.NoError(t, err)

		notebooks, err := srv.ListFromUser(ctx, model.ListNotebookFromUserDTO{
			UserID: userID,
		})
		require.NoError(t, err)

		total := len(notebooks)

		err = srv.Delete(ctx, model.DeleteNotebookDTO{
			UserID:     userID,
			NotebookID: created.NotebookID,
		})
		require.NoError(t, err)

		notebooks, err = srv.ListFromUser(ctx, model.ListNotebookFromUserDTO{
			UserID: userID,
		})
		require.NoError(t, err)
		newTotal := len(notebooks)

		t.Run("it should have 1 less item after delete", func(t *testing.T) {
			assert.Equal(t, total-1, newTotal)
		})
	})

	t.Run("Update Notebook", func(t *testing.T) {
		ctx := context.TODO()
		inputCreate := model.CreateNotebookDTO{
			UserID:      userID,
			Icon:        "folder",
			Name:        "Test Notebook",
			Image:       "google.com/image",
			Description: "description notebook",
		}

		created, err := srv.Create(ctx, inputCreate)
		require.NoError(t, err)

		inputUpdate := model.UpdateNotebookDTO{
			UserID:     userID,
			NotebookID: created.NotebookID,
			Name:       "Updated Notebook",
			Image:      "google.com/image2",
		}
		updated, err := srv.Update(ctx, inputUpdate)
		require.NoError(t, err)

		t.Run("it have same ids", func(t *testing.T) {
			assert.Equal(t, created.NotebookID, updated.NotebookID)
		})

		t.Run("it have same description and icon", func(t *testing.T) {
			assert.Equal(t, created.Description, updated.Description)
			assert.Equal(t, created.Icon, updated.Icon)
		})

		// needs parse on date
		// t.Run("it have same created at and diff updated at", func(t *testing.T) {
		// 	assert.Equal(t, created.CreatedAt, updated.CreatedAt)
		// 	assert.NotEqual(t, created.UpdatedAt, updated.UpdatedAt)
		// })

		t.Run("it have diff name and image", func(t *testing.T) {
			assert.NotEqual(t, created.Name, updated.Name)
			assert.NotEqual(t, created.Image, updated.Image)
		})

		t.Run("it should have the name and image from input update", func(t *testing.T) {
			assert.Equal(t, inputUpdate.Name, updated.Name)
			assert.Equal(t, inputUpdate.Image, updated.Image)
		})
	})
}
