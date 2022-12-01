package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/osmait/cqrs/events"
	"github.com/osmait/cqrs/models"
	"github.com/osmait/cqrs/repository"
	"github.com/osmait/cqrs/search"
)

func onCreateFeed(m events.CreateFeedMessage) {
	feed := models.Feed{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		CreateAt:    m.CreateAt,
	}
	if err := search.IndexFeed(context.Background(), feed); err != nil {
		log.Printf("failed to index feed: %v", err)
	}

}

func listFeedsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	feeds, err := repository.ListFeeds(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		http.Error(w, "query is required", http.StatusBadRequest)
		return
	}
	feeds, err := search.SearchFeed(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)

}
