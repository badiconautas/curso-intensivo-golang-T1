package model

type CreateNotebookInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NotebookEntityDomain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
