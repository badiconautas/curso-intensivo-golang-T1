package repository

import (
	"context"
	"fmt"

	kivik "github.com/go-kivik/kivik/v4"
	"github.com/lucasbadico/intensivo-first-service/model"
)

type NotebookCouch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rev         string `json:"_rev,omitempty"`
}

type Repository struct {
	couchdb *kivik.DB
}

func New(c *kivik.DB) *Repository {
	return &Repository{couchdb: c}
}

func (repo *Repository) Create(ctx context.Context, n model.NotebookEntityDomain) error {
	notebook_couch := NotebookCouch{
		ID:          n.ID,
		Name:        n.Name,
		Description: n.Description,
		Rev:         "",
	}
	resp, err := repo.couchdb.Put(ctx, n.ID, notebook_couch)
	fmt.Printf("couch resp: %+v\n", resp)
	fmt.Printf("couch err: %+v\n", err)
	if err != nil {
		return err
	}

	return nil
}
