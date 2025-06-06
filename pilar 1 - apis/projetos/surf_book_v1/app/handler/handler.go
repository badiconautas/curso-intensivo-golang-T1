package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	// service *service.Service
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func New(
// srv *service.Service,
) *Handler {
	return &Handler{
		// service: srv,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok", Message: "Everthing is cool here..."}); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
	}
}

func (h *Handler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/health", h.Health).Methods("GET")
}
