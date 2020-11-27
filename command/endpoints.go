package command

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/reallocf/Poggle/helper"
	"github.com/segmentio/ksuid"
)

func makeChangeNameEndpoint(svc CommandService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(changeNameRequest)
		id, err := ksuid.Parse(req.Id)
		if err != nil {
			return changeNameResponse{helper.InvalidKUIDError}, nil
		}

		svc.Commander.ChangeName(id, req.NewName)

		return changeNameResponse{""}, nil
	}
}

func makeCreateEndpoint(svc CommandService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)

		id := svc.Commander.Create(req.Name)

		return createResponse{id.String(), ""}, nil
	}
}

func makeDeleteEndpoint(svc CommandService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequest)
		id, err := ksuid.Parse(req.Id)
		if err != nil {
			return deleteResponse{helper.InvalidKUIDError}, nil
		}

		svc.Commander.Delete(id)

		return deleteResponse{""}, nil
	}
}

func makeTurnOffEndpoint(svc CommandService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(turnOffRequest)
		id, err := ksuid.Parse(req.Id)
		if err != nil {
			return turnOffResponse{helper.InvalidKUIDError}, nil
		}

		svc.Commander.TurnOff(id)

		return turnOffResponse{""}, nil
	}
}

func makeTurnOnEndpoint(svc CommandService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(turnOnRequest)
		id, err := ksuid.Parse(req.Id)
		if err != nil {
			return turnOnResponse{helper.InvalidKUIDError}, nil
		}

		svc.Commander.TurnOn(id)

		return turnOnResponse{""}, nil
	}
}
