package handlers

import (
	"encoding/json"
	"net/http"

	"go-leveldb/storage"

	"github.com/go-chi/chi/v5"
)

type LaunchHandler struct {
	store *storage.LaunchStore
}

func NewLaunchHandler(store *storage.LaunchStore) *LaunchHandler {
	return &LaunchHandler{
		store: store,
	}
}

func (h *LaunchHandler) GetAllLaunches(w http.ResponseWriter, r *http.Request) {
	launches, err := h.store.GetAllLaunches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(launches)
}

func (h *LaunchHandler) GetLaunchByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "launch ID is required", http.StatusBadRequest)
		return
	}

	launch, err := h.store.GetLaunch(id)
	if err != nil {
		if err == storage.ErrLaunchNotFound {
			http.Error(w, "launch not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(launch)
}
