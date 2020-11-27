package query

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeGetAllOnRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getAllOnRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeIsOnRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request isOnRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
