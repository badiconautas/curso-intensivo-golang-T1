package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lucasbadico/surfbook-api-v1/model"
)

//CREATE
func (h *Handler) createTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	var input model.CreateTagDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	input.UserID = userID

	tag, err := h.tagSrv.Create(ctx, input)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}

//LIST
func (h *Handler) listTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	input := model.ListTagDTO{
		UserID: userID,
	}

	tags, err := h.TagSrv.ListFromUser(ctx, input)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
	return
}

//DELETE
func (h *Handler) deleteTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	userID, err := uuid.Parse(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, "Header user-id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	TagID, err := uuid.Parse(vars["tag_id"])
	if err != nil {
		http.Error(w, "Header tag_id inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.tagSrv.Delete(ctx, model.DeleteTagDTO{
		TagID: tagID,
		UserID: userID,
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