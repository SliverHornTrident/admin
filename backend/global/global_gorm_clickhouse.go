//go:build (gorm || gorms) && clickhouse

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Clickhouse       *gorm.DB
	ClickhouseConfig config.Gorm
)
