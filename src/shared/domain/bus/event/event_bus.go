package event

type EventBus interface {
	Publish(events ...DomainEvent)
}
