package response

import (
	"github.com/lann/builder"
)

type HealthDetail struct {
	Status      HealthStatus `json:"status" binding:"required"`
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description,omitempty" binding:"required"`
	Error       string       `json:"error,omitempty"`
}

type healthDetailBuilder builder.Builder

func (b healthDetailBuilder) Status(value HealthStatus) healthDetailBuilder {
	return builder.Set(b, "Status", value).(healthDetailBuilder)
}

func (b healthDetailBuilder) Name(value string) healthDetailBuilder {
	return builder.Set(b, "Name", value).(healthDetailBuilder)
}

func (b healthDetailBuilder) Description(value string) healthDetailBuilder {
	return builder.Set(b, "Description", value).(healthDetailBuilder)
}

func (b healthDetailBuilder) Error(value string) healthDetailBuilder {
	return builder.Set(b, "Error", value).(healthDetailBuilder)
}

func (b healthDetailBuilder) Build() HealthDetail {
	return builder.GetStruct(b).(HealthDetail)
}

var HealthDetailBuilder = builder.Register(healthDetailBuilder{}, HealthDetail{}).(healthDetailBuilder)
