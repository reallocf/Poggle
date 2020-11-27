package query

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/reallocf/Poggle/helper"
	"github.com/segmentio/ksuid"
)

func makeIsOnEndpoint(svc QueryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(isOnRequest)
		id, parseErr := ksuid.Parse(req.Id)
		if parseErr != nil {
			return isOnResponse{false, helper.InvalidKUIDError}, nil
		}

		isOn, err := svc.Queryer.IsOn(id)

		if err != nil {
			return isOnResponse{false, err.Error()}, nil
		}
		return isOnResponse{isOn, ""}, nil
	}
}

func makeGetAllOnEndpoint(svc QueryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		allOnKUIDs := svc.Queryer.GetAllOn()

		allOn := helper.ConvertKSUIDsToStrings(allOnKUIDs)

		return getAllOnResponse{allOn, ""}, nil
	}
}
