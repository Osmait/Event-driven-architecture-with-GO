package events

import (
	"context"

	"github.com/osmait/cqrs/models"
)

type EventStore interface {
	Close()
	PublishCreatedFeed(ctx context.Context, feed *models.Feed) error
	SubscribeCreatedFeed(ctx context.Context) (<-chan CreateFeedMessage, error)
	OnCreateFeed(f func(CreateFeedMessage)) error
}

var eventStore EventStore

func Close() {
	eventStore.Close()
}
func PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishCreatedFeed(ctx, feed)

}
func SubscribeCreatedFeed(ctx context.Context) (<-chan CreateFeedMessage, error) {
	return eventStore.SubscribeCreatedFeed(ctx)
}
func OnCreateFeed(ctx context.Context, f func(CreateFeedMessage)) error {
	return eventStore.OnCreateFeed(f)
}
