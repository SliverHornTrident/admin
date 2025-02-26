//go:build gorm || gorms

package interfaces

import (
	"github.com/SliverHornTrident/shadow/config"
	"github.com/SliverHornTrident/shadow/constant"
	"gorm.io/gorm"
)

type Gorms interface {
	Gorm(config config.Gorm) (*gorm.DB, error)
	GormsType() constant.GormsType
}
