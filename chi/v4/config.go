package gichi

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	StatusRoute                = "gi.chi.route.status"
	HealthRoute                = "gi.chi.route.health"
	MiddlewareLogEnabled       = "gi.chi.middleware.log.enabled"
	MiddlewareRecoverEnabled   = "gi.chi.middleware.recover.enabled"
	MiddlewareRealIPEnabled    = "gi.chi.middleware.realip.enabled"
	MiddlewareRequestIDEnabled = "gi.chi.middleware.requestid.enabled"
	MiddlewareNewTID           = "gi.chi.middleware.newtid.enabled"
	MiddlewareNewRelic         = "gi.chi.middleware.newrelic.enabled"
	NewRelicWebResponseEnabled = "gi.chi.middleware.newrelic.webResponseEnabled"
)

func init() {

	log.Println("getting configurations for chi")

	giconfig.Add(StatusRoute, "/resource-status", "define status url")
	giconfig.Add(HealthRoute, "/health", "define health url")
	giconfig.Add(MiddlewareLogEnabled, true, "enable/disable logging request middleware")
	giconfig.Add(MiddlewareRecoverEnabled, true, "enable/disable recover middleware")
	giconfig.Add(MiddlewareRealIPEnabled, true, "enable/disable real ip middleware")
	giconfig.Add(MiddlewareRequestIDEnabled, true, "enable/disable request id middleware")
	giconfig.Add(MiddlewareNewTID, true, "enable/disable new tid middleware")
	giconfig.Add(MiddlewareNewRelic, false, "enable/disable new relic middleware")
	giconfig.Add(NewRelicWebResponseEnabled, true, "enable/disable WebResponse from middleware")
}

func GetStatusRoute() string {
	return giconfig.String(StatusRoute)
}

func GetHealthRoute() string {
	return giconfig.String(HealthRoute)
}

func GetMiddlewareRecoverEnabled() bool {
	return giconfig.Bool(MiddlewareRecoverEnabled)
}
func GetMiddlewareLoggerEnabled() bool {
	return giconfig.Bool(MiddlewareLogEnabled)
}
func GetMiddlewareRealIPEnabled() bool {
	return giconfig.Bool(MiddlewareRealIPEnabled)
}
func GetMiddlewareRequestIDEnabled() bool {
	return giconfig.Bool(MiddlewareRequestIDEnabled)
}

func GetMiddlewareNewTidEnabled() bool {
	return giconfig.Bool(MiddlewareNewTID)
}

func GetMiddlewareNewRelicEnabled() bool {
	return giconfig.Bool(MiddlewareNewRelic)
}

func GetNewRelicWebResponseEnabled() bool {
	return giconfig.Bool(NewRelicWebResponseEnabled)
}
