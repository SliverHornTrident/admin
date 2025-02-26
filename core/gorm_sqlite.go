//go:build (gorm || gorms) && sqlite

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/config"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"strings"
)

var _ interfaces.Corer = (*_sqlite)(nil)

var Sqlite = new(_sqlite)

type _sqlite struct{}

func (c *_sqlite) Gorm(config config.Gorm) (*gorm.DB, error) {
	dialector := sqlite.Open(config.Sqlite())
	db, err := gorm.Open(dialector, config.GormConfig())
	if err != nil {
		return nil, errors.Wrapf(err, "链接失败!")
	}
	sql, err := db.DB()
	if err != nil {
		return nil, errors.Wrapf(err, "获取数据库连接失败!")
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
			source := sqlite.Open(config.Resolvers[i].Sources[j].Sqlite())
			sources = append(sources, source)
		}
		replicas := make([]gorm.Dialector, 0, len(config.Resolvers[i].Replicas))
		for j := 0; j < len(config.Resolvers[i].Replicas); j++ {
			if config.Resolvers[i].Replicas[j].IsEmpty() {
				continue
			}
			replica := sqlite.Open(config.Resolvers[i].Replicas[j].Sqlite())
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

// Init 初始化
func (c *_sqlite) Init() (db *gorm.DB, err error) {
	return db, nil
}

func (c *_sqlite) Name() string {
	return "[shadow][core][gorm][sqlite]"
}

func (c *_sqlite) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Sqlite", &global.SqliteConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_sqlite) IsPanic() bool {
	return true
}

func (c *_sqlite) ConfigName() string {
	return strings.Join([]string{"gorm", "sqlite", gin.Mode(), "yaml"}, ".")
}

func (c *_sqlite) Initialization(ctx context.Context) error {
	db, err := c.Gorm(global.SqliteConfig)
	if err != nil {
		return err
	}
	Gorm.Copy(&global.Sqlite, db)
	return nil
}
