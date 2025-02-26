//go:build (gorm || gorms) && sqlite

package config

import "fmt"

func (c Gorm) Sqlite() string {
	if c.Dbname != "" && c.Config == "file:%s?mode=memory&cache=shared" {
		return fmt.Sprintf(c.Config, c.Dbname)
	}
	if c.Dbname == "" && c.Config == "file::memory:?cache=shared" {
		return c.Config
	}
	return fmt.Sprintf("%s", c.Dbname)
}
