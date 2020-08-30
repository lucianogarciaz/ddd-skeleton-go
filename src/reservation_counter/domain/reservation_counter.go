package domain

import value_object "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/value-object"

type ReservationCounter struct {
	reservationID value_object.ID
	count         int
}

func (r ReservationCounter) Count() int {
	return r.count
}

func (r ReservationCounter) ReservationID() value_object.ID {
	return r.reservationID
}
