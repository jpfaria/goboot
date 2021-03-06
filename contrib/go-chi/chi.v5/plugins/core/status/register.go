package status

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/goignite/v2/core/rest/response"
)

func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	statusRoute := getRoute()

	logger.Tracef("configuring status router on %s in chi", statusRoute)

	statusHandler := NewResourceStatusHandler()

	return &chi.Config{
		Routes: []chi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: statusHandler.Get(),
				Pattern:     statusRoute,
			},
		},
	}, nil
}

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get() http.HandlerFunc {
	resourceStatus := response.NewResourceStatus()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(resourceStatus)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(reqBodyBytes.Bytes())
		w.WriteHeader(http.StatusOK)
	}
}
