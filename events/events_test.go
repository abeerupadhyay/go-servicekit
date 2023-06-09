package events

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Z struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

type DummySub struct {
	eventName string
	sentAt    time.Time
	z         *Z
}

func (s *DummySub) Call(ctx context.Context, d *EventData) {
	s.eventName = string(d.Event)
	s.sentAt = d.SentAt
	s.z = &Z{}
	if err := json.Unmarshal(d.Data, s.z); err != nil {
		panic(err)
	}
}

func TestPublishEvent(t *testing.T) {

	assert := assert.New(t)
	ctx := context.Background()

	event := Register("testing")
	sub := &DummySub{}

	// Register subscriber to the event
	event.Subscribe(sub)

	// Data being sent on publishing event
	z := &Z{Foo: "foo", Bar: 1}

	err := event.Publish(ctx, z)
	if err != nil {
		t.Fatalf("failed to publish event. %v", err)
	}

	assert.Equal(string(event), sub.eventName)
	assert.Equal(z.Foo, sub.z.Foo)
	assert.Equal(z.Bar, sub.z.Bar)
	assert.False(sub.sentAt.IsZero())
}

func TestPublishEventNoSubscriber(t *testing.T) {

	ctx := context.Background()
	event := Register("testing")

	// Data being sent on publishing event
	z := &Z{Foo: "foo", Bar: 1}

	err := event.Publish(ctx, z)
	if err != nil {
		t.Fatalf("failed to publish event. %v", err)
	}
}
