//go:build gorms && tidb

package core

import (
	"github.com/SliverHornTrident/shadow/constant"
)

func (c *_tidb) GormsType() constant.GormsType {
	return constant.GormsTypeTidb
}
