package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lucasbadico/surfbook-api-v1/model"
)

func (h *Handler) createNotebook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	var input model.CreateNotebookDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input.UserID = userID

	notebook, err := h.notebookSrv.Create(ctx, input)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(notebook); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}

func (h *Handler) listNotebook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	input := model.ListNotebookFromUserDTO{
		UserID: userID,
	}

	notebooks, err := h.notebookSrv.ListFromUser(ctx, input)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(notebooks); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}

func (h *Handler) findNotebook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	notebookID, err := uuid.Parse(vars["notebook_id"])
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	notebook, err := h.notebookSrv.FindNotebookFromUser(ctx, model.FindNotebookFromUserDTO{
		NotebookID: notebookID,
		UserID:     userID,
	})

	if err != nil {
		http.Error(w, "Notebook não encontrado: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(notebook); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return

}

func (h *Handler) deleteNotebook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	notebookID, err := uuid.Parse(vars["notebook_id"])
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.notebookSrv.Delete(ctx, model.DeleteNotebookDTO{
		NotebookID: notebookID,
		UserID:     userID,
	})
	if err != nil {
		http.Error(w, "Delete Não executado: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Operação realizada com sucesso",
	})
	if err != nil {
		http.Error(w, "Operação realizada e Erro interno: "+err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *Handler) updateNotebook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	notebookID, err := uuid.Parse(vars["notebook_id"])
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	var input model.UpdateNotebookDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input.UserID = userID
	input.NotebookID = notebookID

	notebook, err := h.notebookSrv.Update(ctx, input)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(notebook); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return

}

func (h *Handler) mountNotebookRoutes(r *mux.Router) {
	r.HandleFunc("/notebooks", h.listNotebook).Methods("GET")
	r.HandleFunc("/notebooks", h.createNotebook).Methods("POST")
	r.HandleFunc("/notebooks/{notebook_id}", h.findNotebook).Methods("GET")
	r.HandleFunc("/notebooks/{notebook_id}", h.updateNotebook).Methods("PATCH")
	r.HandleFunc("/notebooks/{notebook_id}", h.deleteNotebook).Methods("DELETE")
}
