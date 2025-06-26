package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lucasbadico/surfbook-api-v1/service"
)

type Handler struct {
	notebookSrv *service.NotebookService
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok", Message: "Everthing is cool here..."}); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
	}
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func New(
	notebookSrv *service.NotebookService,
) *Handler {
	return &Handler{
		notebookSrv: notebookSrv,
	}
}

func (h *Handler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/health", h.health).Methods("GET")

	h.mountNotebookRoutes(r)

}
