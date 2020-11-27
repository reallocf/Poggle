package toggle

import "github.com/segmentio/ksuid"

type ToggleCommander struct {
	EventStore ToggleEventStore
}

func (t ToggleCommander) ChangeName(id ksuid.KSUID, name string) {
	data := map[string]interface{}{
		"newName": name,
	}

	event := ToggleEvent{
		Id:   id,
		Type: ChangeName,
		Data: data,
	}

	t.EventStore.ExecuteToggleEvent(event)
}

func (t ToggleCommander) Create(name string) ksuid.KSUID {
	id := ksuid.New()

	data := map[string]interface{}{
		"name": name,
		"isEnabled": false,
	}

	event := ToggleEvent{
		Id:   id,
		Type: Create,
		Data: data,
	}

	t.EventStore.ExecuteToggleEvent(event)

	return id
}

func (t ToggleCommander) Delete(id ksuid.KSUID) {
	event := ToggleEvent{
		Id:   id,
		Type: Delete,
	}

	t.EventStore.ExecuteToggleEvent(event)
}

func (t ToggleCommander) TurnOff(id ksuid.KSUID) {
	event := ToggleEvent{
		Id:   id,
		Type: TurnOff,
	}

	t.EventStore.ExecuteToggleEvent(event)
}

func (t ToggleCommander) TurnOn(id ksuid.KSUID) {
	event := ToggleEvent{
		Id:   id,
		Type: TurnOn,
	}

	t.EventStore.ExecuteToggleEvent(event)
}
