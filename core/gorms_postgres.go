//go:build gorms && tidb

package core

import (
	"github.com/SliverHornTrident/shadow/constant"
)

func (c *_postgres) GormsType() constant.GormsType {
	return constant.GormsTypePostgres
}
