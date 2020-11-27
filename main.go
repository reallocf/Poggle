package main

import (
	"github.com/reallocf/Poggle/command"
	"github.com/reallocf/Poggle/event_store"
	"github.com/reallocf/Poggle/query"
	"github.com/reallocf/Poggle/toggle"
	"github.com/segmentio/ksuid"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	eventStore := event_store.InMemoryEventStore{make(map[ksuid.KSUID][]toggle.ToggleEvent)}

	wg.Add(2)
	go func() {
		defer wg.Done()
		command.RunServer(8081, eventStore)
	}()
	go func() {
		defer wg.Done()
		query.RunServer(8080, eventStore)
	}()

	wg.Wait()
}

