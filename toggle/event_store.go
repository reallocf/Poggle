package toggle

import "github.com/segmentio/ksuid"

type ToggleEventStore interface {
	ExecuteToggleEvent(event ToggleEvent)
	ConstructToggleAggregate(id ksuid.KSUID) (*ToggleAggregate, error)
	ConstructAllToggleAggregates() []ToggleAggregate
}
