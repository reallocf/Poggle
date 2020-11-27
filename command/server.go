package command

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/reallocf/Poggle/helper"
	"github.com/reallocf/Poggle/toggle"
	"log"
	"net/http"
	"strconv"
)

type CommandService struct {
	Commander toggle.ToggleCommander
}

func RunServer(port int, eventStore toggle.ToggleEventStore) {
	server := http.NewServeMux()
	svc := CommandService{toggle.ToggleCommander{eventStore}}

	changeNameHandler := httptransport.NewServer(
		makeChangeNameEndpoint(svc),
		decodeChangeNameRequest,
		helper.EncodeResponse,
	)

	createHandler := httptransport.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		helper.EncodeResponse,
	)

	deleteHandler := httptransport.NewServer(
		makeDeleteEndpoint(svc),
		decodeDeleteRequest,
		helper.EncodeResponse,
	)

	turnOffHandler := httptransport.NewServer(
		makeTurnOffEndpoint(svc),
		decodeTurnOffRequest,
		helper.EncodeResponse,
	)

	turnOnHandler := httptransport.NewServer(
		makeTurnOnEndpoint(svc),
		decodeTurnOnRequest,
		helper.EncodeResponse,
	)

	server.Handle("/changeName", changeNameHandler)
	server.Handle("/create", createHandler)
	server.Handle("/delete", deleteHandler)
	server.Handle("/turnOff", turnOffHandler)
	server.Handle("/turnOn", turnOnHandler)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), server))
}
