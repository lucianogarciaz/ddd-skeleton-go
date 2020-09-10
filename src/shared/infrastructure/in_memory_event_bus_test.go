package infrastructure_test

import (
	"errors"
	"testing"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/infrastructure"

	errors2 "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/application/errors"

	"github.com/stretchr/testify/require"
)

func TestEventBus(t *testing.T) {
	eventName := "moq"
	domainEventMoq := &domain.DomainEventMock{
		EventNameFunc: func() string {
			return eventName
		},
	}
	domainEventSubscriberMoq := &domain.DomainEventSubscriberMock{
		HandleFunc: func(domain.DomainEvent) error {
			return nil
		},
		SubscribedToFunc: func() []domain.DomainEvent {
			return []domain.DomainEvent{domainEventMoq}
		},
	}

	t.Run(`Given an in memory event bus and a domain event with no subscribers,
	when it's called,
	then it returns nil`, func(t *testing.T) {
		subscribers := map[string][]domain.DomainEventSubscriber{
			"random.event": {domainEventSubscriberMoq},
		}
		eventBus := infrastructure.NewInMemoryEventBus(subscribers)
		err := eventBus.Publish(domainEventMoq)
		require.NoError(t, err)
		require.Empty(t, domainEventSubscriberMoq.HandleCalls())
	})

	t.Run(`Given an in memory event bus and a domain event subscriber that errors,
	when it's called,
	then it return a multi error`, func(t *testing.T) {
		returnedError := errors.New("error")
		domainEventSubscriberErrorsMoq := &domain.DomainEventSubscriberMock{
			HandleFunc: func(domain.DomainEvent) error {
				return returnedError
			},
			SubscribedToFunc: func() []domain.DomainEvent {
				return []domain.DomainEvent{domainEventMoq}
			},
		}
		subscribers := map[string][]domain.DomainEventSubscriber{
			(domainEventMoq).EventName(): {domainEventSubscriberErrorsMoq},
		}
		eventBus := infrastructure.NewInMemoryEventBus(subscribers)
		err := eventBus.Publish(domainEventMoq)
		require.True(t, err.(*errors2.MultiError).Is(returnedError))
	})

	t.Run(`Given an in memory event bus and a domain event,
	when it's called,
	then it consume the same message correctly`, func(t *testing.T) {
		subscribers := map[string][]domain.DomainEventSubscriber{
			(domainEventMoq).EventName(): {domainEventSubscriberMoq},
		}
		eventBus := infrastructure.NewInMemoryEventBus(subscribers)
		err := eventBus.Publish(domainEventMoq)
		require.NoError(t, err)
		require.Len(t, domainEventSubscriberMoq.HandleCalls(), 1)
	})
}
