package mongodb

import (
	"context"
	"strings"

	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func NewClient(ctx context.Context, o *Options) (client *mongo.Client, database *mongo.Database, err error) {

	l := log.FromContext(ctx)

	nrMon := nrmongo.NewCommandMonitor(nil)

	clientOptions := options.Client().ApplyURI(o.Uri)

	if o.NewRelic.Enabled {
		clientOptions.SetMonitor(nrMon)
	}

	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	connFields, err := connstring.Parse(o.Uri)

	if err != nil {
		return nil, nil, err
	}

	database = client.Database(connFields.Database)

	l.Infof("Connected to MongoDB server: %v", strings.Join(connFields.Hosts, ","))

	if o.Health.Enabled {
		configureHealthCheck(client, o)
	}

	return client, database, err
}

func NewDefaultClient(ctx context.Context) (*mongo.Client, *mongo.Database, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}

func configureHealthCheck(client *mongo.Client, o *Options) {
	mc := NewClientChecker(client)
	hc := health.NewHealthChecker("mongodb", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
