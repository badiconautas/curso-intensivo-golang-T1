package model 

import(
	"time"

	"github.com/google/uuid"
)

type UserTagEntity struct {
	TagID 		uuid.UUID `json:"tag_id"`
	UserID      uuid.UUID `json:"user_id"`
	TagName        string `json:"tag_name"`
	Color 		   string `json:"color"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type CreateTagDTO struct {
	TagID 		uuid.UUID `json:"tag_id"`
	UserID 		 uuid.UUID 
	TagName 		string `json:"tag_name"`
	Color 			string `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserTagDTO struct {
	TagID 		uuid.UUID `json:"tag_id"`
	UserID 		 uuid.UUID 
	TagName 		string `json:"tag_name"`
	Color 			string `json:"color"`
}

type ListTagDTO struct{
	UserID     uuid.UUID
}

type DeleteTagDTO struct {
	TagID	   uuid.UUID
	UserID     uuid.UUID
}