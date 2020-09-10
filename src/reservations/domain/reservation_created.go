package domain

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

var _ domain.DomainEvent = ReservationCreatedDomainEvent{}

const reservationCreatedDomainEvent = "reservation.created"

type ReservationCreatedDomainEvent struct {
	domain.BasicDomainEvent
	hotel   string
	barcode string
}

func NewReservationCreatedDomainEvent(id string, hotel string, barcode string, eventID *string, occuredOn *string) ReservationCreatedDomainEvent {
	domainEvent := domain.NewDomainEvent(id, eventID, occuredOn)
	return ReservationCreatedDomainEvent{
		BasicDomainEvent: domainEvent,
		hotel:            hotel,
		barcode:          barcode,
	}
}

func (ReservationCreatedDomainEvent) EventName() string {
	return reservationCreatedDomainEvent
}

func (ReservationCreatedDomainEvent) FromPrimitives(aggregateID string, body map[string]interface{}, eventID string, occurredOn string) (domain.DomainEvent, error) {
	hotel, ok := body["hotel"].(string)
	if !ok {
		return domain.BasicDomainEvent{}, domain.ErrExtractFromPrimitives
	}
	barcode, ok := body["barcode"].(string)
	if !ok {
		return domain.BasicDomainEvent{}, domain.ErrExtractFromPrimitives
	}
	return NewReservationCreatedDomainEvent(aggregateID, hotel, barcode, &eventID, &occurredOn), nil
}

func (r ReservationCreatedDomainEvent) ToPrimitives() (map[string]interface{}, error) {
	return map[string]interface{}{
		"hotel":   r.hotel,
		"barcode": r.barcode,
	}, nil
}
