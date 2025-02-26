//go:build (gorm || gorms) && mysql

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Mysql       *gorm.DB
	MysqlConfig config.Gorm
)
