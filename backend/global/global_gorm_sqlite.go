//go:build (gorm || gorms) && sqlite

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Sqlite       *gorm.DB
	SqliteConfig config.Gorm
)
