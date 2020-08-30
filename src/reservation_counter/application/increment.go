package application

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/reservations/domain"
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/bus/event"
)

var _ event.DomainEventSubscriber = IncrementReservationCounterOnReservationCreated{}

type IncrementReservationCounterOnReservationCreated struct {
}

func (i IncrementReservationCounterOnReservationCreated) SubscribedTo() []event.DomainEvent {
	return []event.DomainEvent{domain.ReservationCreatedDomainEvent{}}
}

func (i IncrementReservationCounterOnReservationCreated) Handle(event.DomainEvent) error {
	return nil
}
