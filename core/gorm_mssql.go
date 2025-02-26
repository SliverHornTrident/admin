//go:build (gorm || gorms) && (mssql || sqlserver)

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/config"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"strings"
)

var Mssql = new(mssql)

type mssql struct{}

func (c *mssql) Gorm(config config.Gorm) (*gorm.DB, error) {
	dialector := sqlserver.New(config.Mssql())
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
			source := sqlserver.New(config.Resolvers[i].Sources[j].Mssql())
			sources = append(sources, source)
		}
		replicas := make([]gorm.Dialector, 0, len(config.Resolvers[i].Replicas))
		for j := 0; j < len(config.Resolvers[i].Replicas); j++ {
			if config.Resolvers[i].Replicas[j].IsEmpty() {
				continue
			}
			replica := sqlserver.New(config.Resolvers[i].Replicas[j].Mssql())
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

func (c *mssql) Name() string {
	return "[shadow][core][gorm][mssql]"
}

func (c *mssql) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Mssql", &global.MssqlConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *mssql) IsPanic() bool {
	return true
}

func (c *mssql) ConfigName() string {
	return strings.Join([]string{"gorm", "mssql", gin.Mode(), "yaml"}, ".")
}

func (c *mssql) Initialization(ctx context.Context) error {
	db, err := c.Gorm(global.MssqlConfig)
	if err != nil {
		return err
	}
	Gorm.Copy(&global.Mssql, db)
	return nil
}
