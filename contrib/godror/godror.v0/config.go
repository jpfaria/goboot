package godror

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root            = "gi.godror"
	dataSourceName  = root + ".dataSourceName"
	connMaxLifetime = root + ".connMaxLifetime"
	maxIdleConns    = root + ".maxIdleConns"
	maxOpenConns    = root + ".maxOpenConns"
	PluginsRoot     = root + ".plugins"
)

func init() {

	config.Add(dataSourceName, "", "database name and connection information")
	config.Add(connMaxLifetime, 0*time.Second, "sets the maximum amount of time a connection may be reused. If d <= 0, connections are reused forever")
	config.Add(maxIdleConns, 2, "sets the maximum number of connections in the idle connection pool.")
	config.Add(maxOpenConns, 5, "sets the maximum number of open connections to the database.")
}
