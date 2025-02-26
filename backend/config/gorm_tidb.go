//go:build tidb

package config

import (
	"fmt"
	"gorm.io/driver/mysql"
)

func (c Gorm) Tidb() mysql.Config {
	if c.Config == "" {
		c.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Config)
	return mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: true,
	}
}
