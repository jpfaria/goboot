package datadog

import (
	"context"

	datadog "github.com/b2wdigital/goignite/v2/contrib/datadog/dd-trace-go.v1"
	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1"
	"github.com/b2wdigital/goignite/v2/core/log"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

func Register(ctx context.Context, conn *mongo.Conn) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating mongo in datadog")

	conn.ClientOptions.SetMonitor(mongotrace.NewMonitor(mongotrace.WithServiceName(datadog.Service())))

	logger.Debug("mongo successfully integrated in datadog")

	return nil
}
