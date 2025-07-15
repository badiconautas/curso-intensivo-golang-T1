package repository

import (
	"context"
	"database/sql"
	"embed"

	"github.com/google/uuid"
	"github.com/lucasbadico/surfbook-api-v1/model"
)
//banco
type UserTagsRepository struct {
	db *sql.DB
}

//go:embed _queries/tag_create.sql
var createTagQuery string

//go:embed _queries/tag_list.sql
var listTagQuery string

//go:embed _queries/tag_delete.sql
var deleteTagQuery string

//cria um novo repositorio e injeta a conexão do banco
func NewUserTagRepository(db *sql.DB) *UserTagsRepository {
	return &UserTagsRepository{
		db : db,
	}
}
//CREATE
func (repo *UserTagsRepository) CreateTag(ctx context.Context, tag *model.UserTagEntity) error {
	_, err := repo.db.ExecContext(
		ctx,
		createTagQuery,
		tag.TagID,
		tag.UserID,
		tag.TagName,
		tag.Color,
		tag.CreatedAt,
		tag.DeletedAt
	)
	return err
}

//LIST
func (repo *UserTagsRepository) ListTag(ctx context.Context, userID uuid.UUID) ([]*model.UserTagEntity, error) {
	rows, err := repo.db.QueryContext(ctx, listTagQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*model.UserTagEntity

	for rows.Next() {
		var t model.UserTagEntity
		if err := rows.Scan(
			&t.TagID,
			&t.UserID,
			&t.TagName,
			&t.Color,
			&t.Created_at,
			&t.Deleted_at
		); err != nil {
			return nil, err
		}
		tags = append(tags, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

//DELETE
func (repo *UserTagsRepository) SoftDeleteTag(ctx context.Context, deleteIt model.DeleteTagDTO) error {
	_, err := repo.db.Exec(
		deleteTagQuery,
		deleteIt.TagID,
		deleteIt.UserID,
	)
	return err
}