package rscpgx

import (
	"time"
)

type PlatformConfig struct {
	DSN               string        `config:"dsn,secret,required"`
	ConnectionTimeout time.Duration `config:"connection_timeout"`
}
