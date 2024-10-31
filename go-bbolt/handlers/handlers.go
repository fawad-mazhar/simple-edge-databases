package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"go-bbolt/db"
	"go-bbolt/models"

	"github.com/go-chi/chi/v5"
)

type PageData struct {
	Buckets       []string
	CurrentBucket string
	Data          map[string]string
}

type Handler struct {
	db *db.Database
}

func NewHandler(db *db.Database) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetAllLaunches(w http.ResponseWriter, r *http.Request) {
	// var launches []models.Launch
	bucketName := "launches"
	launches, err := h.db.GetAllLaunches(bucketName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(launches)
}

func (h *Handler) GetLaunchByID(w http.ResponseWriter, r *http.Request) {
	launchID := chi.URLParam(r, "id")
	var launch models.Launch

	err := h.db.GetLaunchByID("launches", launchID, &launch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(launch)
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) Buckets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buckets := h.db.ListBuckets()
	templates := template.Must(template.ParseGlob("ui/templates/*.html"))
	templates.ExecuteTemplate(w, "buckets.html", models.BucketPage{Buckets: buckets})
}

func (h *Handler) Bucket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	bucketName := chi.URLParam(r, "name")
	pairs := h.db.GetBucketKeys(bucketName)
	templates := template.Must(template.ParseGlob("ui/templates/*.html"))
	templates.ExecuteTemplate(w, "bucket.html", models.KeyValuePage{
		BucketName: bucketName,
		Pairs:      pairs,
	})
}

func (h *Handler) KeyValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	bucketName := chi.URLParam(r, "bucket")
	key := chi.URLParam(r, "key")
	response := h.db.GetKeyValueInJson(bucketName, key)
	json.NewEncoder(w).Encode(response)
}
