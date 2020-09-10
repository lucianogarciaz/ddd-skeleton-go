package application

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/reservations/domain"
	domain2 "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

var _ domain2.DomainEventSubscriber = IncrementReservationCounterOnReservationCreated{}

type IncrementReservationCounterOnReservationCreated struct {
}

func (i IncrementReservationCounterOnReservationCreated) SubscribedTo() []domain2.DomainEvent {
	return []domain2.DomainEvent{domain.ReservationCreatedDomainEvent{}}
}

func (i IncrementReservationCounterOnReservationCreated) Handle(domain2.DomainEvent) error {
	return nil
}
