package query

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/reallocf/Poggle/helper"
	"github.com/reallocf/Poggle/toggle"
	"log"
	"net/http"
	"strconv"
)

type QueryService struct {
	Queryer toggle.ToggleQueryer
}

func RunServer(port int, eventStore toggle.ToggleEventStore) {
	server := http.NewServeMux()
	querySvc := QueryService{toggle.ToggleQueryer{eventStore}}

	getAllOnHandler := httptransport.NewServer(
		makeGetAllOnEndpoint(querySvc),
		decodeGetAllOnRequest,
		helper.EncodeResponse,
	)

	isOnHandler := httptransport.NewServer(
		makeIsOnEndpoint(querySvc),
		decodeIsOnRequest,
		helper.EncodeResponse,
	)

	server.Handle("/getAllOn", getAllOnHandler)
	server.Handle("/isOn", isOnHandler)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), server))
}
