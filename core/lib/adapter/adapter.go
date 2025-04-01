package adapter

import "github.com/nonamecat19/go-orm/core/lib/config"

type Adapter interface {
	GetConnString(config config.ORMConfig) string
	GetDbDriver() string
}
