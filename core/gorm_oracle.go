//go:build (gorm || gorms) && oracle

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/config"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/dzwvip/oracle"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"strings"
)

var Oracle = new(_oracle)

type _oracle struct{}

func (c *_oracle) Gorm(config config.Gorm) (*gorm.DB, error) {
	dialector := oracle.New(config.Oracle())
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
			source := oracle.New(config.Resolvers[i].Sources[j].Oracle())
			sources = append(sources, source)
		}
		replicas := make([]gorm.Dialector, 0, len(config.Resolvers[i].Replicas))
		for j := 0; j < len(config.Resolvers[i].Replicas); j++ {
			if config.Resolvers[i].Replicas[j].IsEmpty() {
				continue
			}
			replica := oracle.New(config.Resolvers[i].Replicas[j].Oracle())
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

func (c *_oracle) Name() string {
	return "[shadow][core][gorm][oracle]"
}

func (c *_oracle) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Oracle", &global.OracleConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_oracle) IsPanic() bool {
	return true
}

func (c *_oracle) ConfigName() string {
	return strings.Join([]string{"gorm", "oracle", gin.Mode(), "yaml"}, ".")
}

func (c *_oracle) Initialization(ctx context.Context) error {
	db, err := c.Gorm(global.OracleConfig)
	if err != nil {
		return err
	}
	Gorm.Copy(&global.Oracle, db)
	return nil
}
