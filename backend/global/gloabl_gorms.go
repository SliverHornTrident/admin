//go:build gorms && (tidb || mysql || postgres || sqlite || clickhouse)

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Gorms       map[string]*gorm.DB
	GormsConfig config.Gorms
)
