//go:build gorms && (tidb || mysql || postgres || sqlite || clickhouse)

package config

import "github.com/SliverHornTrident/shadow/constant"

type Gorms []GormsChildren

type GormsChildren struct {
	Type   constant.GormsType `json:"Type" yaml:"Type" mapstructure:"Type"`       // 数据库类型
	Name   string             `json:"Name" yaml:"Name" mapstructure:"Name"`       // 业务数据库名
	Enable bool               `json:"Enable" yaml:"Enable" mapstructure:"Enable"` // 是否开启
	Gorm   `json:",inline" yaml:",inline" mapstructure:",squash"`
}
