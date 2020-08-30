package aggregate_root

import "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/bus/event"

type AggregateRoot interface {
	PullDomainEvents() []event.DomainEvent
	Record(event.DomainEvent)
}

type BasicAggregateRoot struct {
	domainEvents []event.DomainEvent
}

func (a *BasicAggregateRoot) PullDomainEvents() []event.DomainEvent {
	domainEvents := a.domainEvents
	a.domainEvents = nil
	return domainEvents
}

func (a *BasicAggregateRoot) Record(domainEvent event.DomainEvent) {
	a.domainEvents = append(a.domainEvents, domainEvent)
}
