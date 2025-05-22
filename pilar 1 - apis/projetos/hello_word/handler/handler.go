package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kivik/kivik/v4"
	"github.com/gorilla/mux"
)

type Handler struct {
	couchdb *kivik.DB
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Notebook struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func New(c *kivik.DB) *Handler {
	return &Handler{
		couchdb: c,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok", Message: "Everthing is cool here..."}); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
	}
}

// Create
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	var nb Notebook
	if err := json.NewDecoder(r.Body).Decode(&nb); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.couchdb.Put(ctx, nb.ID, nb)
	if err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}

	println("couch resp", resp)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}

// Update

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

// List
