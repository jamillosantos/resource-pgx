package rscpgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Resource struct {
	*pgxpool.Pool
	name                      string
	config                    PlatformConfig
	disableConnectionOnStatup bool
}

func NewResource(config PlatformConfig, options ...Option) *Resource {
	opt := defaultOpts()
	for _, option := range options {
		option(&opt)
	}
	return &Resource{
		name:   opt.name,
		config: config,

		disableConnectionOnStatup: opt.disableConnectionOnStatup,
	}
}

func (r *Resource) Name() string {
	return r.name
}

func (r *Resource) Start(ctx context.Context) error {
	conn, err := pgxpool.New(ctx, r.config.DSN)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	r.Pool = conn
	if r.disableConnectionOnStatup {
		return nil
	}
	return conn.Ping(ctx)
}

func (r *Resource) Stop(ctx context.Context) error {
	r.Pool.Close()
	return nil
}
