//go:build redis

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"github.com/go-redis/redis/v8"
)

var (
	Redis        *redis.Client
	RedisConfig  config.Redis
	RedisCluster *redis.ClusterClient
)
