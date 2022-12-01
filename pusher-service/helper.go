package main

import "time"

type CreateFeedMessage struct {
	Type        string    `json:"type"`
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}

func newCreateFeedMessage(id, title, description string, createdAt time.Time) *CreateFeedMessage {
	return &CreateFeedMessage{
		Type:        "created_feed",
		ID:          id,
		Title:       title,
		Description: description,
		CreateAt:    createdAt,
	}
}
