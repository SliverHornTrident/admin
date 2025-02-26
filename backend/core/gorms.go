//go:build gorms && (tidb || mysql || postgres || sqlite || clickhouse)

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/config"
	"github.com/SliverHornTrident/shadow/constant"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
)

var _ interfaces.Corer = (*Gorms)(nil)

type Gorms struct {
	realizes map[constant.GormsType]interfaces.Gorms
}
type GormsOptions func(config config.GormsChildren) (*gorm.DB, error)

func NewGorms(options ...interfaces.Gorms) *Gorms {
	gorms := &Gorms{}
	length := len(options)
	gorms.realizes = make(map[constant.GormsType]interfaces.Gorms, length)
	for i := 0; i < len(options); i++ {
		gorms.realizes[options[i].GormsType()] = options[i]
	}
	return gorms
}

func (c *Gorms) Name() string {
	return "[shadow][core][gorms]"
}

func (*Gorms) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Gorms", &global.GormsConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *Gorms) IsPanic() bool {
	return false
}

func (c *Gorms) ConfigName() string {
	return strings.Join([]string{"gorms", gin.Mode(), "yaml"}, ".")
}

func (c *Gorms) Initialization(ctx context.Context) error {
	for i := 0; i < len(global.GormsConfig); i++ {
		if !global.GormsConfig[i].Enable {
			continue
		}
		value, ok := c.realizes[global.GormsConfig[i].Type]
		if !ok {
			return errors.New("gorms type not found")
		}
		db, err := value.Gorm(global.GormsConfig[i].Gorm)
		if err != nil {
			return err
		}
		_, ok = global.Gorms[global.GormsConfig[i].Name]
		if ok {
			return errors.New("name 重复!")
		}
		global.Gorms[global.GormsConfig[i].Name] = db
	}
	return nil
}
