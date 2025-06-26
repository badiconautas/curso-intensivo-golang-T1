package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lucasbadico/surfbook-api-v1/model"
	"github.com/lucasbadico/surfbook-api-v1/repository"
)

type NotebookService struct {
	repo *repository.NotebookRepository
}

func (srv *NotebookService) Create(ctx context.Context, input model.CreateNotebookDTO) (*model.NotebookEntity, error) {
	now := time.Now()
	entity := &model.NotebookEntity{
		NotebookID:  uuid.New(),
		UserID:      input.UserID,
		Icon:        input.Icon,
		Description: input.Description,
		Name:        input.Name,
		Image:       input.Image,
		DeletedAt:   nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err := srv.repo.CreateNotebook(ctx, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (srv *NotebookService) FindNotebookFromUser(ctx context.Context, input model.FindNotebookFromUserDTO) (*model.NotebookEntity, error) {
	notebook, err := srv.repo.FindNotebookByUserIDandID(ctx, input.UserID, input.NotebookID)
	if err != nil {
		return nil, err
	}

	return notebook, nil
}

func (srv *NotebookService) ListFromUser(ctx context.Context, input model.ListNotebookFromUserDTO) ([]*model.NotebookEntity, error) {
	return srv.repo.ListNotebooks(ctx, input.UserID)
}

func (srv *NotebookService) Delete(ctx context.Context, input model.DeleteNotebookDTO) error {
	return srv.repo.SoftDeleteNotebook(ctx, input)
}

func (srv *NotebookService) Update(ctx context.Context, input model.UpdateNotebookDTO) (*model.NotebookEntity, error) {
	err := srv.repo.UpdateNotebook(ctx, input)
	if err != nil {
		return nil, err
	}

	notebook, err := srv.repo.FindNotebookByUserIDandID(ctx, input.UserID, input.NotebookID)
	if err != nil {
		return nil, err
	}

	return notebook, nil
}

func NewNotebookService(
	repo *repository.NotebookRepository,
) *NotebookService {
	return &NotebookService{
		repo: repo,
	}
}
