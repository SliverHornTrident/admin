//go:build gorms && clickhouse

package core

import (
	"github.com/SliverHornTrident/shadow/constant"
)

func (c *_clickhouse) GormsType() constant.GormsType {
	return constant.GormsTypeClickhouse
}
