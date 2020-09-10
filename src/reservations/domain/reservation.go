package domain

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

// Reservation is the aggregate root
type Reservation struct {
	domain.BasicAggregateRoot
	id      domain.ID
	hotel   string
	barcode string
}

// NewReservation is a constructor
func NewReservation(id domain.ID, hotel string, barcode string) Reservation {
	return Reservation{
		id:      id,
		hotel:   hotel,
		barcode: barcode,
	}
}

// CreateReservation creates a reservation and records the event.
func CreateReservation(id domain.ID, hotel string, barcode string) Reservation {
	reservation := NewReservation(id, hotel, barcode)
	reservation.Record(NewReservationCreatedDomainEvent(id.String(), hotel, barcode, nil, nil))
	return reservation
}
