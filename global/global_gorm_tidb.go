//go:build (gorm || gorms) && tidb

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Tidb       *gorm.DB
	TidbConfig config.Gorm
)
