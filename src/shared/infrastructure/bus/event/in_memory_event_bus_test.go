package event_test

import (
	"errors"
	"testing"

	errors2 "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/application/errors"

	event2 "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/infrastructure/bus/event"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/bus/event"
)

func TestSpace(t *testing.T) {
	eventName := "moq"
	domainEventMoq := &event.DomainEventMock{
		EventNameFunc: func() string {
			return eventName
		},
	}
	domainEventSubscriberMoq := &event.DomainEventSubscriberMock{
		HandleFunc: func(event.DomainEvent) error {
			return nil
		},
		SubscribedToFunc: func() []event.DomainEvent {
			return []event.DomainEvent{domainEventMoq}
		},
	}

	t.Run(`Given an in memory event bus and a domain event with no subscribers,
	when it's called,
	then it returns nil`, func(t *testing.T) {
		subscribers := map[string][]event.DomainEventSubscriber{
			"random.event": {domainEventSubscriberMoq},
		}
		eventBus := event2.NewInMemoryEventBus(subscribers)
		err := eventBus.Publish(domainEventMoq)
		require.NoError(t, err)
		require.Empty(t, domainEventSubscriberMoq.HandleCalls())
	})

	t.Run(`Given an in memory event bus and a domain event subscriber that errors,
	when it's called,
	then it return a multi error`, func(t *testing.T) {
		returnedError := errors.New("error")
		domainEventSubscriberErrorsMoq := &event.DomainEventSubscriberMock{
			HandleFunc: func(event.DomainEvent) error {
				return returnedError
			},
			SubscribedToFunc: func() []event.DomainEvent {
				return []event.DomainEvent{domainEventMoq}
			},
		}
		subscribers := map[string][]event.DomainEventSubscriber{
			(domainEventMoq).EventName(): {domainEventSubscriberErrorsMoq},
		}
		eventBus := event2.NewInMemoryEventBus(subscribers)
		err := eventBus.Publish(domainEventMoq)
		require.True(t, err.(*errors2.MultiError).Is(returnedError))
	})

	t.Run(`Given an in memory event bus and a domain event,
	when it's called,
	then it consume the same message correctly`, func(t *testing.T) {
		subscribers := map[string][]event.DomainEventSubscriber{
			(domainEventMoq).EventName(): {domainEventSubscriberMoq},
		}
		eventBus := event2.NewInMemoryEventBus(subscribers)
		err := eventBus.Publish(domainEventMoq)
		require.NoError(t, err)
		require.Len(t, domainEventSubscriberMoq.HandleCalls(), 1)
	})
}
