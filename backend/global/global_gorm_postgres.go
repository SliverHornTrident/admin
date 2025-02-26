//go:build (gorm || gorms) && postgres

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Postgres       *gorm.DB
	PostgresConfig config.Gorm
)
