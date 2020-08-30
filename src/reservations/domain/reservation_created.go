package domain

import "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/bus/event"

var _ event.DomainEvent = ReservationCreatedDomainEvent{}

const reservationCreatedDomainEvent = "reservation.created"

type ReservationCreatedDomainEvent struct {
	event.BasicDomainEvent
	hotel   string
	barcode string
}

func NewReservationCreatedDomainEvent(id string, hotel string, barcode string, eventID *string, occuredOn *string) ReservationCreatedDomainEvent {
	domainEvent := event.NewDomainEvent(id, eventID, occuredOn)
	return ReservationCreatedDomainEvent{
		BasicDomainEvent: domainEvent,
		hotel:            hotel,
		barcode:          barcode,
	}
}

func (ReservationCreatedDomainEvent) EventName() string {
	return reservationCreatedDomainEvent
}

func (ReservationCreatedDomainEvent) FromPrimitives(aggregateID string, body map[string]interface{}, eventID string, occurredOn string) (event.DomainEvent, error) {
	hotel, ok := body["hotel"].(string)
	if !ok {
		return event.BasicDomainEvent{}, event.ErrExtractFromPrimitives
	}
	barcode, ok := body["barcode"].(string)
	if !ok {
		return event.BasicDomainEvent{}, event.ErrExtractFromPrimitives
	}
	return NewReservationCreatedDomainEvent(aggregateID, hotel, barcode, &eventID, &occurredOn), nil
}

func (r ReservationCreatedDomainEvent) ToPrimitives() (map[string]interface{}, error) {
	return map[string]interface{}{
		"hotel":   r.hotel,
		"barcode": r.barcode,
	}, nil
}
