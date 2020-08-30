package event

type DomainEventSubscriber interface {
	SubscribedTo() []DomainEvent
	Handle(DomainEvent) error
}
