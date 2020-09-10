package controllers

import (
	"fmt"
	"net/http"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/reservations/application"
)

type ReservationController struct {
	commandBus domain.CommandBus
}

func NewReservationController(cmd domain.CommandBus) ReservationController {
	return ReservationController{
		commandBus: cmd,
	}
}

func (r ReservationController) Handle(http.ResponseWriter, *http.Request) {
	r.ReservationPostController("some hotel", "some barcode")
}

func (r ReservationController) ReservationPostController(hotel string, barcode string) {
	cmd := application.NewCreateReservationCommand(hotel, barcode)
	if err := r.commandBus.Dispatch(cmd); err != nil {
		fmt.Println(err)
	}
}
