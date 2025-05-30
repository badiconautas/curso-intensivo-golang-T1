package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucasbadico/intensivo-first-service/model"
	"github.com/lucasbadico/intensivo-first-service/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		repo: r,
	}
}

// Create
func (srv *Service) Create(ctx context.Context, input model.CreateNotebookInputDTO) (*model.NotebookEntityDomain, error) {
	id := uuid.NewString()
	err := validateNotebookName(input.Name)
	if err != nil {
		return nil, err
	}
	// & eu quero o valor do endereco de memoria
	new_notebook := &model.NotebookEntityDomain{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	// * eu quero o valor que esta guardano nesse endereco de memoria
	// err := srv.repo.Create(ctx, *new_notebook)

	err = srv.repo.Create(ctx, new_notebook)
	if err != nil {
		return nil, err
	}

	return new_notebook, nil
}

// // Update
// func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	ctx := context.TODO()

// 	vars := mux.Vars(r)
// 	docID := vars["notebook_id"]

// 	var existingNB Notebook
// 	err := h.couchdb.Get(ctx, vars["notebook_id"]).ScanDoc(&existingNB)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Erro ao buscar Notebook: %v", err), http.StatusNotFound)
// 		return
// 	}

// 	rev := existingNB.Rev // Aqui pega a revisão do documento
// 	if rev == "" {
// 		http.Error(w, "Falha ao obter a rev", http.StatusInternalServerError)
// 		return
// 	}

// 	var updated Notebook
// 	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
// 		http.Error(w, "Erro ao decodificar JSON de entrada", http.StatusBadRequest)
// 		return
// 	}

// 	updated.ID = docID
// 	updated.Rev = rev

// 	newRev, err := h.couchdb.Put(ctx, updated.ID, updated)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Erro ao atualizar documento: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	updated.Rev = newRev
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(updated)
// }

// // Get
func (srv *Service) Get(ctx context.Context, id string) (*model.NotebookEntityDomain, error) {
	nb, err := srv.repo.Get(ctx, id)

	if err != nil {
		return nil, err
	}
	return nb, nil
}

// Delete
func (srv *Service) Delete(ctx context.Context, id string) error {

	err := srv.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// // List
// func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	ctx := context.TODO()

// 	rows := h.couchdb.AllDocs(ctx, kivik.Params(map[string]interface{}{"include_docs": true}))

// 	if rows != nil {
// 		http.Error(w, fmt.Sprintf("Erro ao listar documentos: %v", rows), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var notebooks []Notebook

// 	for rows.Next() {
// 		var nb Notebook
// 		if err := rows.ScanDoc(&nb); err != nil {
// 			http.Error(w, fmt.Sprintf("Erro ao ler documento: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		id, err := rows.ID()
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Erro ao obter ID do documento: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		nb.ID = id
// 		fmt.Printf("couch resp: %+v\n", nb)

// 		notebooks = append(notebooks, nb)
// 	}

// 	if err := rows.Err(); err != nil {
// 		http.Error(w, fmt.Sprintf("Erro durante iteração: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	if err := json.NewEncoder(w).Encode(notebooks); err != nil {
// 		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
// 		return
// 	}
// }
