package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/osmait/cqrs/events"
	"github.com/osmait/cqrs/models"
	"github.com/osmait/cqrs/repository"
	"github.com/segmentio/ksuid"
)

type createdFeedRequest struct {
	Title       string `json:"title"`
	Description string `jsom:"description"`
}

func createFeedHandler(w http.ResponseWriter, r *http.Request) {
	var req createdFeedRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	feed := models.Feed{
		ID:          id.String(),
		Title:       req.Title,
		Description: req.Description,
		CreateAt:    createdAt,
	}

	if err := repository.InsertFeed(r.Context(), &feed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := events.PublishCreatedFeed(r.Context(), &feed); err != nil {
		log.Printf("failed to publish created feed event: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}
