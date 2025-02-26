//go:build (gorm || gorms) && (mssql || sqlserver)

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Mssql       *gorm.DB
	MssqlConfig config.Gorm
)
