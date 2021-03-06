package infrastructure

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/application/errors"
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

type InMemoryEventBus struct {
	subscribers map[string][]domain.DomainEventSubscriber
}

func NewInMemoryEventBus(subscribers map[string][]domain.DomainEventSubscriber) InMemoryEventBus {
	return InMemoryEventBus{subscribers}
}

func (im InMemoryEventBus) Publish(events ...domain.DomainEvent) error {
	multierror := errors.NewMultiError()
	for _, e := range events {
		es, ok := im.subscribers[e.EventName()]
		if !ok {
			continue
		}
		for _, s := range es {
			if err := s.Handle(e); err != nil {
				multierror.Add(err)
			}
		}
	}
	return multierror.ErrResult()
}
