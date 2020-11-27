package event_store

import (
	"errors"
	"github.com/reallocf/Poggle/helper"
	"github.com/reallocf/Poggle/toggle"
	"github.com/segmentio/ksuid"
)

type InMemoryEventStore struct {
	Store map[ksuid.KSUID][]toggle.ToggleEvent
}

func (es InMemoryEventStore) ExecuteToggleEvent(event toggle.ToggleEvent) {
	es.Store[event.Id] = append(es.Store[event.Id], event)
}

func (es InMemoryEventStore) ConstructToggleAggregate(id ksuid.KSUID) (*toggle.ToggleAggregate, error) {
	var aggregate *toggle.ToggleAggregate
	events := es.Store[id]
	for _, event := range(events) {
		aggregate = toggle.FoldEvent(aggregate, event)
	}
	if aggregate == nil {
		return nil, errors.New(helper.MissingKUIDError)
	} else {
		return aggregate, nil
	}
}

func (es InMemoryEventStore) ConstructAllToggleAggregates() []toggle.ToggleAggregate {
	var aggregates []toggle.ToggleAggregate
	for id, _ := range(es.Store) {
		aggregate, err := es.ConstructToggleAggregate(id)
		if err == nil {
			aggregates = append(aggregates, *aggregate)
		}
	}
	return aggregates
}

