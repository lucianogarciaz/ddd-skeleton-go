package main

import (
	"fmt"
	"net/http"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/infrastructure"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/reservations/application"

	"github.com/lucianogarciaz/ddd-skeleton-go/app/src/controllers"
)

func main() {
	bus := infrastructure.NewInMemoryCommandBus()
	reservationHandler := application.NewCreateReservationCommandHandler(application.NewReservationCreator(infrastructure.InMemoryEventBus{}))
	if err := bus.RegisterHandler(application.CreateReservationCommand{}, reservationHandler); err != nil {
		fmt.Println("Error trying to register a handler: ", err)
		return
	}

	r := controllers.NewReservationController(&bus)
	http.HandleFunc("/", r.Handle)
	http.ListenAndServe(":8080", nil)
}
