//go:build gorms && mysql

package core

import (
	"github.com/SliverHornTrident/shadow/constant"
)

func (c *_mysql) GormsType() constant.GormsType {
	return constant.GormsTypeMysql
}
