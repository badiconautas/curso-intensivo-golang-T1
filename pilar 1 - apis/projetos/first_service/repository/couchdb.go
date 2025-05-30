package repository

import (
	"context"
	"errors"

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

func (repo *Repository) Create(ctx context.Context, n *model.NotebookEntityDomain) error {
	notebook_couch := NotebookCouch{
		ID:          n.ID,
		Name:        n.Name,
		Description: n.Description,
		Rev:         "",
	}
	_, err := repo.couchdb.Put(ctx, n.ID, notebook_couch)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Delete(ctx context.Context, id string) error {
	var nb NotebookCouch

	row := repo.couchdb.Get(ctx, id)
	if err := row.ScanDoc(&nb); err != nil {
		return err
	}

	rev := nb.Rev // Aqui pega a revisão do documento
	if rev == "" {
		return errors.New("Revisao não encontrada")
	}

	_, err := repo.couchdb.Delete(ctx, id, rev)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) Get(ctx context.Context, id string) (*model.NotebookEntityDomain, error) {
	var nb NotebookCouch

	err := repo.couchdb.Get(ctx, id).ScanDoc(&nb)
	if err != nil {
		return nil, err
	}

	nd := &model.NotebookEntityDomain{
		ID:          nb.ID,
		Name:        nb.Name,
		Description: nb.Description,
	}
	return nd, nil
}
