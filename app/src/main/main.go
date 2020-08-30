package main

import (
	"fmt"
	"net/http"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/infrastructure/bus/event"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/reservations/application"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/infrastructure/bus/command"

	"github.com/lucianogarciaz/ddd-skeleton-go/app/src/controllers"
)

func main() {
	bus := command.NewInMemoryCommandBus()
	reservationHandler := application.NewCreateReservationCommandHandler(application.NewReservationCreator(event.InMemoryEventBus{}))
	if err := bus.RegisterHandler(application.CreateReservationCommand{}, reservationHandler); err != nil {
		fmt.Println("Error trying to register a handler: ", err)
		return
	}

	r := controllers.NewReservationController(&bus)
	http.HandleFunc("/", r.Handle)
	http.ListenAndServe(":8080", nil)
}
