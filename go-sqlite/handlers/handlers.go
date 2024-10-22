package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"go-sqlite/db"
)

type Handler struct {
	db *db.Database
}

func NewHandler(db *db.Database) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetAllLaunches(w http.ResponseWriter, r *http.Request) {
	launches, err := h.db.GetAllLaunches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(launches)
}

func (h *Handler) GetLaunchByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	launch, err := h.db.GetLaunchByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if launch == nil {
		http.Error(w, "Launch not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(launch)
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
