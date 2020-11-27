package toggle

import "github.com/segmentio/ksuid"

type ToggleEventType int

const (
	ChangeName ToggleEventType = iota
	Create
	Delete
	TurnOff
	TurnOn
)

type ToggleEvent struct {
	Id   ksuid.KSUID
	Type ToggleEventType
	Data map[string]interface{}
}
