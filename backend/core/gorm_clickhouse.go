//go:build (gorm || gorms) && clickhouse

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/config"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"strings"
)

var _ interfaces.Corer = (*_clickhouse)(nil)

var Clickhouse = new(_clickhouse)

type _clickhouse struct{}

func (c *_clickhouse) Gorm(config config.Gorm) (*gorm.DB, error) {
	dialector := clickhouse.New(config.Clickhouse())
	db, err := gorm.Open(dialector, config.GormConfig())
	if err != nil {
		return nil, errors.Wrap(err, "链接失败!")
	}
	sql, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "获取数据库连接失败!")
	}
	sql.SetMaxIdleConns(config.MaxIdleCones)
	sql.SetMaxOpenConns(config.MaxOpenCones)
	sql.SetConnMaxLifetime(config.ConnMaxLifetime)
	sql.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	resolver := new(dbresolver.DBResolver)
	length := len(config.Resolvers)
	for i := 0; i < length; i++ {
		sources := make([]gorm.Dialector, 0, len(config.Resolvers[i].Sources))
		for j := 0; j < len(config.Resolvers[i].Sources); j++ {
			if config.Resolvers[i].Sources[j].IsEmpty() {
				continue
			}
			source := clickhouse.New(config.Resolvers[i].Sources[j].Clickhouse())
			sources = append(sources, source)
		}
		replicas := make([]gorm.Dialector, 0, len(config.Resolvers[i].Replicas))
		for j := 0; j < len(config.Resolvers[i].Replicas); j++ {
			if config.Resolvers[i].Replicas[j].IsEmpty() {
				continue
			}
			replica := clickhouse.New(config.Resolvers[i].Replicas[j].Clickhouse())
			replicas = append(replicas, replica)
		}
		resolver.Register(dbresolver.Config{
			Sources:           sources,
			Replicas:          replicas,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: config.Resolvers[i].TraceResolverMode,
		}, config.Resolvers[i].Data()...).
			SetConnMaxLifetime(config.Resolvers[i].ConnMaxLifetime).
			SetConnMaxIdleTime(config.Resolvers[i].ConnMaxIdleTime).
			SetMaxIdleConns(config.Resolvers[i].MaxIdleCones).
			SetMaxOpenConns(config.Resolvers[i].MaxOpenCones)
	}
	err = db.Use(resolver)
	if err != nil {
		return nil, errors.Wrap(err, "注册负载均衡失败!")
	}
	return db, nil
}

func (c *_clickhouse) Name() string {
	return "[shadow][core][gorm][clickhouse]"
}

func (c *_clickhouse) IsPanic() bool {
	return true
}

func (c *_clickhouse) ConfigName() string {
	return strings.Join([]string{"gorm", "clickhouse", gin.Mode(), "yaml"}, ".")
}

func (c *_clickhouse) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Clickhouse", &global.ClickhouseConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_clickhouse) Initialization(ctx context.Context) error {
	db, err := c.Gorm(global.ClickhouseConfig)
	if err != nil {
		return err
	}
	Gorm.Copy(&global.Clickhouse, db)
	return nil
}
