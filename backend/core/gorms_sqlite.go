//go:build gorms && sqlite

package core

import (
	"github.com/SliverHornTrident/shadow/constant"
)

func (c *_sqlite) GormsType() constant.GormsType {
	return constant.GormsTypeSqlite
}
