//go:build (gorm || gorms) && oracle

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"gorm.io/gorm"
)

var (
	Oracle       *gorm.DB
	OracleConfig config.Gorm
)
