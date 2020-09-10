package domain

import "errors"

type DomainEvent interface {
	FromPrimitives(aggregateID string, body map[string]interface{}, eventID string, occurredOn string) (DomainEvent, error)
	ToPrimitives() (map[string]interface{}, error)
	EventName() string
	AggregateID() string
	EventID() *string
	OcurredOn() *string
}

type BasicDomainEvent struct {
	DomainEvent
	aggreagateID string
	eventID      *string
	occurredOn   *string
}

func NewDomainEvent(aggregateID string, eventID *string, occuredOn *string) BasicDomainEvent {
	return BasicDomainEvent{
		aggreagateID: aggregateID,
		eventID:      eventID,
		occurredOn:   occuredOn,
	}
}

func (d BasicDomainEvent) AggregateID() string {
	return d.aggreagateID
}

func (d BasicDomainEvent) EventID() *string {
	return d.eventID
}

func (d BasicDomainEvent) OcurredOn() *string {
	return d.occurredOn
}

var ErrExtractFromPrimitives = errors.New("domain event from primitive can not be extracted")
