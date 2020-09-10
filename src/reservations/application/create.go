package application

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/reservations/domain"
	domain2 "github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

var _ domain2.Command = &CreateReservationCommand{}

// CreateReservationCommandName is self-described
const CreateReservationCommandName = "createReservation"

// CreateReservationCommand  is a command
type CreateReservationCommand struct {
	hotel   string
	barcode string
}

// NewCreateReservationCommand is a constructor
func NewCreateReservationCommand(hotel string, barcode string) CreateReservationCommand {
	return CreateReservationCommand{
		hotel:   hotel,
		barcode: barcode,
	}
}

// Name is a getter
func (crc CreateReservationCommand) Name() string {
	return CreateReservationCommandName
}

// Hotel is a getter
func (crc CreateReservationCommand) Hotel() string {
	return crc.hotel
}

// Barcode is a getter
func (crc CreateReservationCommand) Barcode() string {
	return crc.barcode
}

var _ domain2.CommandHandler = CreateReservationCommandHandler{}

// CreateReservationCommandHandler is a command handler
type CreateReservationCommandHandler struct {
	reservationCreator ReservationCreator
}

// NewCreateReservationCommandHandler is a constructor
func NewCreateReservationCommandHandler(reservationCreator ReservationCreator) CreateReservationCommandHandler {
	return CreateReservationCommandHandler{reservationCreator: reservationCreator}
}

// Handle is self-described
func (r CreateReservationCommandHandler) Handle(cmd domain2.Command) error {
	c, ok := cmd.(CreateReservationCommand)
	if !ok {
		return domain2.ErrInvalidCommand
	}
	return r.reservationCreator.Handle(c.Hotel(), c.Barcode())
}

// ReservationCreator is the implementation of the use case
type ReservationCreator struct {
	eventBus domain2.EventBus
}

// NewReservationCreator is a constructor
func NewReservationCreator(eventBus domain2.EventBus) ReservationCreator {
	return ReservationCreator{
		eventBus: eventBus,
	}
}

// Handle is self-described
func (rc ReservationCreator) Handle(hotel string, barcode string) error {
	reservation := domain.CreateReservation(domain2.NewID(), hotel, barcode)
	rc.eventBus.Publish(reservation.PullDomainEvents()...)
	return nil
}
