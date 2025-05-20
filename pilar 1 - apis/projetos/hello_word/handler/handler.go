package handler

import (
	"net/http"

	"github.com/go-kivik/kivik/v4"
)

type Handler struct {
	couchdb *kivik.DB
}

func New(c *kivik.DB) *Handler {
	return &Handler{
		couchdb: c,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
