package event

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/bus/event"
)

type InMemoryEventBus struct{}

func (InMemoryEventBus) Publish(events ...event.DomainEvent) {
}
