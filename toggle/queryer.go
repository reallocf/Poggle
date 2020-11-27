package toggle

import "github.com/segmentio/ksuid"

type ToggleQueryer struct {
	EventStore ToggleEventStore
}

func (t ToggleQueryer) GetAllOn() []ksuid.KSUID {
	aggregates := t.EventStore.ConstructAllToggleAggregates()
	var res []ksuid.KSUID
	for _, aggregate := range(aggregates) {
		if aggregate.isEnabled {
			res = append(res, aggregate.toggleId)
		}
	}
	return res
}

func (t ToggleQueryer) IsOn(id ksuid.KSUID) (bool, error) {
	aggregate, err := t.EventStore.ConstructToggleAggregate(id)
	if err != nil {
		return false, err
	}
	return aggregate.isEnabled, nil
}
