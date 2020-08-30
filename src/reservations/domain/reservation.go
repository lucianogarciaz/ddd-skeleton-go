package domain

import (
	aggregate_root "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/aggregate_root"
	value_object "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/value-object"
)

// Reservation is the aggregate root
type Reservation struct {
	aggregate_root.BasicAggregateRoot
	id      value_object.ID
	hotel   string
	barcode string
}

// NewReservation is a constructor
func NewReservation(id value_object.ID, hotel string, barcode string) Reservation {
	return Reservation{
		id:      id,
		hotel:   hotel,
		barcode: barcode,
	}
}

// CreateReservation creates a reservation and records the event.
func CreateReservation(id value_object.ID, hotel string, barcode string) Reservation {
	reservation := NewReservation(id, hotel, barcode)
	reservation.Record(NewReservationCreatedDomainEvent(id.String(), hotel, barcode, nil, nil))
	return reservation
}
