package domain

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

type ReservationCounter struct {
	reservationID domain.ID
	count         int
}

func (r ReservationCounter) Count() int {
	return r.count
}

func (r ReservationCounter) ReservationID() domain.ID {
	return r.reservationID
}
