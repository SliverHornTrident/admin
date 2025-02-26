//go:build oracle

package config

import (
	"fmt"
	"github.com/dzwvip/oracle"
)

func (c Gorm) Oracle() oracle.Config {
	if c.Config == "" {
		// TODO 默认配置
		c.Config = ""
	}
	dsn := fmt.Sprintf("oracle://%s:%s@%s:%d/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Config)
	return oracle.Config{
		DSN:               dsn,
		DefaultStringSize: 191, // string 类型字段的默认长度
	}
}
