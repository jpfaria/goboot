package response

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/lann/builder"
)

type Health struct {
	Status  HealthStatus   `json:"status" binding:"required"`
	Details []HealthDetail `json:"details,omitempty" binding:"required"`
}

type healthBuilder builder.Builder

func (b healthBuilder) Status(value HealthStatus) healthBuilder {
	return builder.Set(b, "Status", value).(healthBuilder)
}

func (b healthBuilder) Details(value []HealthDetail) healthBuilder {
	return builder.Set(b, "Details", value).(healthBuilder)
}

func (b healthBuilder) Build() Health {
	return builder.GetStruct(b).(Health)
}

var HealthBuilder = builder.Register(healthBuilder{}, Health{}).(healthBuilder)

func NewHealth(ctx context.Context) (Health, int) {

	var details []HealthDetail

	all := health.CheckAll(ctx)

	httpStatus := http.StatusOK
	healthStatus := Ok

	for _, v := range all {

		healthDetailStatus := Ok

		if !v.IsOk() {
			healthDetailStatus = Down
		}

		var err string

		if v.Error != nil {
			err = v.Error.Error()
		}

		healthDetailResponse := HealthDetailBuilder.
			Name(v.HealthCheck.Name).
			Description(v.HealthCheck.Description).
			Status(healthDetailStatus).
			Error(err).
			Build()

		details = append(details, healthDetailResponse)

		if !v.IsOk() && httpStatus != http.StatusServiceUnavailable {
			if v.HealthCheck.IsRequired() {
				httpStatus = http.StatusServiceUnavailable
				healthStatus = Down
			} else {
				httpStatus = http.StatusMultiStatus
				healthStatus = Partial
			}
		}
	}

	return HealthBuilder.Details(details).Status(healthStatus).Build(), httpStatus
}
