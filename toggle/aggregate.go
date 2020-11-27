package toggle

import "github.com/segmentio/ksuid"

type ToggleAggregate struct {
	toggleId ksuid.KSUID
	name string
	isEnabled bool
}

func FoldEvent(agg *ToggleAggregate, event ToggleEvent) *ToggleAggregate {
	switch event.Type {
	case ChangeName:
		agg.name = event.Data["newName"].(string)
	case Create:
		agg = &ToggleAggregate{
			toggleId: event.Id,
			name: event.Data["name"].(string),
			isEnabled: event.Data["isEnabled"].(bool),
		}
	case Delete:
		agg = nil
	case TurnOn:
		agg.isEnabled = true
	case TurnOff:
		agg.isEnabled = false
	}
	return agg
}