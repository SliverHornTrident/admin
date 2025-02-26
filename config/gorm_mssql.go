//go:build (gorm || gorms) && (mssql || sqlserver)

package config

import (
	"fmt"
	"gorm.io/driver/sqlserver"
)

func (c Gorm) Mssql() sqlserver.Config {
	if c.Config == "" {
		c.Config = "encrypt=disable"
	}
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Config)
	return sqlserver.Config{
		DSN:               dsn,
		DefaultStringSize: 191,
	}
}
