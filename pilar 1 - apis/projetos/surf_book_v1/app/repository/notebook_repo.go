package repository

import (
	"context"
	"database/sql"
	"embed"

	"github.com/google/uuid"
	"github.com/lucasbadico/surfbook-api-v1/model"
)

type NotebookRepository struct {
	db *sql.DB
}

//go:embed _queries/notebook_list.sql
var listNotebookQuery string

//go:embed _queries/notebook_create.sql
var createNotebookQuery string

//go:embed _queries/notebook_delete.sql
var deleteNotebookQuery string

//go:embed _queries/notebook_update.sql
var updateNotebookQuery string

//go:embed _queries/notebook_find_by_user_id_notebook_id.sql
var findByUserIDAndIDNotebookQuery string

var _empty = embed.FS{}

func (repo *NotebookRepository) CreateNotebook(ctx context.Context, notebook *model.NotebookEntity) error {
	_, err := repo.db.Exec(
		createNotebookQuery,
		notebook.NotebookID,
		notebook.UserID,
		notebook.Icon,
		notebook.Name,
		notebook.Description,
		notebook.Image,
		notebook.CreatedAt,
		notebook.UpdatedAt,
		notebook.DeletedAt,
	)
	return err
}

func (repo *NotebookRepository) ListNotebooks(ctx context.Context, userID uuid.UUID) ([]*model.NotebookEntity, error) {
	rows, err := repo.db.QueryContext(ctx, listNotebookQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notebooks []*model.NotebookEntity

	for rows.Next() {
		var nb model.NotebookEntity
		if err := rows.Scan(
			&nb.NotebookID,
			&nb.UserID,
			&nb.Name,
			&nb.Description,
			&nb.Icon,
			&nb.Image,
			&nb.CreatedAt,
			&nb.UpdatedAt,
		); err != nil {
			return nil, err
		}
		notebooks = append(notebooks, &nb)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notebooks, nil
}

func (repo *NotebookRepository) FindNotebookByUserIDandID(ctx context.Context, userID uuid.UUID, notebookID uuid.UUID) (*model.NotebookEntity, error) {
	rows, err := repo.db.QueryContext(ctx, findByUserIDAndIDNotebookQuery, userID, notebookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notebooks []*model.NotebookEntity

	for rows.Next() {
		var nb model.NotebookEntity
		if err := rows.Scan(
			&nb.NotebookID,
			&nb.UserID,
			&nb.Name,
			&nb.Description,
			&nb.Icon,
			&nb.Image,
			&nb.CreatedAt,
			&nb.UpdatedAt,
		); err != nil {
			return nil, err
		}
		notebooks = append(notebooks, &nb)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notebooks[0], nil
}

func (repo *NotebookRepository) SoftDeleteNotebook(ctx context.Context, deleteIt model.DeleteNotebookDTO) error {
	_, err := repo.db.Exec(
		deleteNotebookQuery,
		deleteIt.NotebookID,
		deleteIt.UserID,
	)
	return err
}

func (repo *NotebookRepository) UpdateNotebook(ctx context.Context, updateIt model.UpdateNotebookDTO) error {
	_, err := repo.db.Exec(
		updateNotebookQuery,
		updateIt.NotebookID,  // 1
		updateIt.UserID,      // 2
		updateIt.Name,        // 3
		updateIt.Description, // 4
		updateIt.Icon,        // 5
		updateIt.Image,       // 6
	)
	return err
}

func NewNotebookRepository(db *sql.DB) *NotebookRepository {
	return &NotebookRepository{
		db: db,
	}
}
