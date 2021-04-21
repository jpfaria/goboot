package godror

import (
	"context"
	"database/sql"

	"github.com/b2wdigital/goignite/v2/core/log"
	_ "github.com/godror/godror"
)

type Ext func(context.Context, *sql.DB) error

func NewDBWithOptions(ctx context.Context, o *Options, exts ...Ext) (db *sql.DB, err error) {

	logger := log.FromContext(ctx)

	db, err = sql.Open("godror", o.DataSourceName)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	db.SetConnMaxLifetime(o.ConnMaxLifetime)
	db.SetMaxIdleConns(o.MaxIdleConns)
	db.SetMaxOpenConns(o.MaxOpenConns)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	for _, ext := range exts {
		if err := ext(ctx, db); err != nil {
			panic(err)
		}
	}

	logger.Info("Connected to Oracle (godror) server")

	return db, err
}

func NewDB(ctx context.Context, exts ...Ext) (*sql.DB, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewDBWithOptions(ctx, o, exts...)
}
