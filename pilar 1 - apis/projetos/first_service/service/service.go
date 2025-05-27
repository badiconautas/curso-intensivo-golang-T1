package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kivik "github.com/go-kivik/kivik/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
func (srv *Service) Create(ctx context.Context, input CreateNotebookInput) (*Notebook, error) {
	id := uuid.NewString()

	new_notebook := &Notebook{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}
	return new_notebook, nil
}

// Update
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	vars := mux.Vars(r)
	docID := vars["notebook_id"]

	var existingNB Notebook
	err := h.couchdb.Get(ctx, vars["notebook_id"]).ScanDoc(&existingNB)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar Notebook: %v", err), http.StatusNotFound)
		return
	}

	rev := existingNB.Rev // Aqui pega a revisão do documento
	if rev == "" {
		http.Error(w, "Falha ao obter a rev", http.StatusInternalServerError)
		return
	}

	var updated Notebook
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Erro ao decodificar JSON de entrada", http.StatusBadRequest)
		return
	}

	updated.ID = docID
	updated.Rev = rev

	newRev, err := h.couchdb.Put(ctx, updated.ID, updated)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao atualizar documento: %v", err), http.StatusInternalServerError)
		return
	}

	updated.Rev = newRev
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

// Get
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	vars := mux.Vars(r)

	var nb Notebook

	err := h.couchdb.Get(ctx, vars["notebook_id"]).ScanDoc(&nb)
	if err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(nb); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

// Delete
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	vars := mux.Vars(r)
	docID := vars["notebook_id"]

	var nb Notebook
	row := h.couchdb.Get(ctx, docID)
	if err := row.ScanDoc(&nb); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar documento: %v", err), http.StatusNotFound)
		return
	}

	rev := nb.Rev // Aqui pega a revisão do documento
	if rev == "" {
		http.Error(w, "Falha ao obter a rev", http.StatusInternalServerError)
		return
	}

	_, err := h.couchdb.Delete(ctx, docID, rev)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao deletar documento: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Sucesso ao deletar o documento",
		"documento": docID,
	})
}

// List
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	rows := h.couchdb.AllDocs(ctx, kivik.Params(map[string]interface{}{"include_docs": true}))

	if rows != nil {
		http.Error(w, fmt.Sprintf("Erro ao listar documentos: %v", rows), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notebooks []Notebook

	for rows.Next() {
		var nb Notebook
		if err := rows.ScanDoc(&nb); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao ler documento: %v", err), http.StatusInternalServerError)
			return
		}
		id, err := rows.ID()
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao obter ID do documento: %v", err), http.StatusInternalServerError)
			return
		}
		nb.ID = id
		fmt.Printf("couch resp: %+v\n", nb)

		notebooks = append(notebooks, nb)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Erro durante iteração: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(notebooks); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}
