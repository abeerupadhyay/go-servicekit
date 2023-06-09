package events

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type EventRegistry map[string][]Subscriber

var (
	// Global registery of all the events
	registry EventRegistry

	// Locks to handle registering and subscribing to events
	rm sync.Mutex
	sm sync.Mutex

	ErrEventNotRegistered = errors.New("event not registered")
)

func init() {
	registry = make(EventRegistry, 0)
	rm = sync.Mutex{}
	sm = sync.Mutex{}
}

type EventData struct {
	Event  Event     // Mostly used for logging
	SentAt time.Time // Timestamp when event was published
	Data   []byte    // Actual data of the event encoded as json
}

type Subscriber interface {
	Call(context.Context, *EventData)
}

type SubscriberFunc func(context.Context, *EventData)

func (f SubscriberFunc) Call(ctx context.Context, e *EventData) {
	f(ctx, e)
}

type Event string

func (e Event) Subscribe(s Subscriber) error {
	return Subscribe(string(e), s)
}

func (e Event) Publish(ctx context.Context, v any) error {
	return Publish(ctx, string(e), v)
}

func Register(name string) Event {
	name = strings.ToLower(name)
	rm.Lock()
	if _, ok := registry[name]; !ok {
		registry[name] = make([]Subscriber, 0)
	}
	rm.Unlock()
	return Event(name)
}

func Subscribe(name string, s Subscriber) error {
	name = strings.ToLower(name)
	sm.Lock()
	if _, ok := registry[name]; !ok {
		return fmt.Errorf("%w - '%s'", ErrEventNotRegistered, name)
	}
	registry[name] = append(registry[name], s)
	sm.Unlock()
	return nil
}

func Publish(ctx context.Context, name string, v any) error {

	name = strings.ToLower(name)
	subs, ok := registry[name]
	if !ok {
		return fmt.Errorf("%w - '%s'", ErrEventNotRegistered, name)
	}

	// Simply ignore if no subscribers are registered
	if len(subs) == 0 {
		return nil
	}

	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	d := &EventData{
		Event:  Event(name),
		SentAt: time.Now().UTC(),
		Data:   bytes,
	}

	var wg sync.WaitGroup
	for _, sub := range subs {
		wg.Add(1)
		go func(ctx context.Context, s Subscriber) {
			defer wg.Done()
			s.Call(ctx, d)
		}(ctx, sub)
	}

	wg.Wait()
	return nil
}
