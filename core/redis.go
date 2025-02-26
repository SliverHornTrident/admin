//go:build redis

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var _ interfaces.Corer = (*_redis)(nil)

var Redis = new(_redis)

type _redis struct{}

func (c *_redis) Name() string {
	return "[shadow][core][redis]"
}

func (c *_redis) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Redis", &global.RedisConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_redis) IsPanic() bool {
	return false
}

func (c *_redis) ConfigName() string {
	return strings.Join([]string{"redis", gin.Mode(), "yaml"}, ".")
}

func (c *_redis) Initialization(ctx context.Context) error {
	if len(global.RedisConfig.Addresses) > 0 {
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    global.RedisConfig.Addresses,
			Password: global.RedisConfig.Password,
		})
		_, err := client.Ping(ctx).Result()
		if err != nil {
			return errors.Wrap(err, "链接redis集群失败!")
		}
		global.RedisCluster = client
	}
	client := redis.NewClient(&redis.Options{
		DB:       global.RedisConfig.Db,
		Addr:     global.RedisConfig.Address,
		Username: global.RedisConfig.Username,
		Password: global.RedisConfig.Password,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return errors.Wrap(err, "链接redis失败!")
	}
	global.Redis = client
	return nil
}
