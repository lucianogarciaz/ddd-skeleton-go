package domain

type DomainEventSubscriber interface {
	SubscribedTo() []DomainEvent
	Handle(DomainEvent) error
}
