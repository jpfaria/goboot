package response

import (
	"github.com/b2wdigital/goignite/v2/core/info"
	"github.com/lann/builder"
)

type ResourceStatusResponse struct {
	ApplicationName       string `json:"applicationName" binding:"required"`
	ImplementationVersion string `json:"implementationVersion" binding:"required"`
	ImplementationBuild   string `json:"implementationBuild" binding:"required"`
	CommitSHA             string `json:"commitSHA" binding:"required"`
	BuildDate             string `json:"buildDate" binding:"required"`
}

type resourceStatusResponseBuilder builder.Builder

func (b resourceStatusResponseBuilder) ApplicationName(applicationName string) resourceStatusResponseBuilder {
	return builder.Set(b, "ApplicationName", applicationName).(resourceStatusResponseBuilder)
}

func (b resourceStatusResponseBuilder) ImplementationVersion(implementationVersion string) resourceStatusResponseBuilder {
	return builder.Set(b, "ImplementationVersion", implementationVersion).(resourceStatusResponseBuilder)
}

func (b resourceStatusResponseBuilder) ImplementationBuild(implementationBuild string) resourceStatusResponseBuilder {
	return builder.Set(b, "ImplementationBuild", implementationBuild).(resourceStatusResponseBuilder)
}

func (b resourceStatusResponseBuilder) CommitSHA(commitSHA string) resourceStatusResponseBuilder {
	return builder.Set(b, "CommitSHA", commitSHA).(resourceStatusResponseBuilder)
}

func (b resourceStatusResponseBuilder) BuildDate(buildDate string) resourceStatusResponseBuilder {
	return builder.Set(b, "BuildDate", buildDate).(resourceStatusResponseBuilder)
}

func (b resourceStatusResponseBuilder) Build() ResourceStatusResponse {
	return builder.GetStruct(b).(ResourceStatusResponse)
}

var ResourceStatusResponseBuilder = builder.Register(resourceStatusResponseBuilder{}, ResourceStatusResponse{}).(resourceStatusResponseBuilder)

func NewResourceStatus() ResourceStatusResponse {

	return ResourceStatusResponseBuilder.
		ApplicationName(info.AppName).
		ImplementationBuild(info.BuildVersion).
		ImplementationVersion(info.Version).
		BuildDate(info.BuildDate).
		CommitSHA(info.CommitSHA).
		Build()

}
