//go:build (gorm || gorms) && postgres

package config

import (
	"fmt"
	"gorm.io/driver/postgres"
)

func (c Gorm) Postgres() postgres.Config {
	if c.Config == "" {
		c.Config = "sslmode=disable TimeZone=Asia/Shanghai"
	}
	dsn := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s %s`, c.Host, c.Port, c.Username, c.Password, c.Dbname, c.Config)
	return postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage: 禁用隐式预准备语句用法
	}
}
