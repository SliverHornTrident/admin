//go:build viper && embed

package core

import (
	"embed"
	"github.com/SliverHornTrident/shadow/core/internal"
	"github.com/SliverHornTrident/shadow/global"
)

func (c *_viper) Embed(configs embed.FS) error {
	global.Configs = configs
	err := internal.Viper.Embed(c.pwd, c.configs)
	if err != nil {
		return err
	}
	return nil
}
