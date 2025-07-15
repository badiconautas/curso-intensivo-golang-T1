package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lucasbadico/surfbook-api-v1/model"
	"github.com/lucasbadico/surfbook-api-v1/repository"
)

type UserTagService struct {
	repo *repository.UserTagsRepository
}

func NewUserTagService(repo *repository.UserTagsRepository) *UserTagService {
	return &UserTagService{
		repo: repo,
	}
}

//CREATE
func (srv *UserTagService) Create(ctx context.Context, input model.CreateTagDTO) (*model.UserTagEntity, error) {
	now := time.Now()
	entity := &model.UserTagEntity{
		TagID:  	 uuid.New(),
		UserID:      input.UserID,
		TagName:	 input.TagName,
		Color: 		 input.Color,
		DeletedAt:   nil,
		CreatedAt:   now,
	}
	err := srv.repo.CreateTag(ctx, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

//LIST
func (srv *UserTagService) ListFromUser(ctx context.Context, input model.ListTagDTO) ([]*model.UserTagEntity, error) {
	return srv.repo.ListTag(ctx, input.UserID)
}

//DELETE
func (srv *UserTagService) Delete(ctx context.Context, input model.DeleteTagDTO) error {
	return srv.repo.SoftDeleteTag(ctx, input)
}