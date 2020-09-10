package domain

type AggregateRoot interface {
	PullDomainEvents() []DomainEvent
	Record(DomainEvent)
}

type BasicAggregateRoot struct {
	domainEvents []DomainEvent
}

func (a *BasicAggregateRoot) PullDomainEvents() []DomainEvent {
	domainEvents := a.domainEvents
	a.domainEvents = nil
	return domainEvents
}

func (a *BasicAggregateRoot) Record(domainEvent DomainEvent) {
	a.domainEvents = append(a.domainEvents, domainEvent)
}
