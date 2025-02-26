//go:build elasticsearch

package core

import (
	"context"
	"fmt"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Elasticsearch = new(_elasticsearch)

type _elasticsearch struct{}

func (c *_elasticsearch) Name() string {
	return "[shadow][core][elasticsearch]"
}

func (c *_elasticsearch) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Elasticsearch", &global.ElasticsearchConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_elasticsearch) IsPanic() bool {
	return true
}

func (c *_elasticsearch) ConfigName() string {
	return strings.Join([]string{"elasticsearch", gin.Mode(), "yaml"}, ".")
}

func (c *_elasticsearch) Initialization(ctx context.Context) error {
	config := global.ElasticsearchConfig.Config()
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		return errors.Wrap(err, "elasticsearch client failed!")
	}
	global.ElasticsearchTypedClient, err = elasticsearch.NewTypedClient(config)
	if err != nil {
		return errors.Wrap(err, "elasticsearch typed client failed!")
	}
	global.Elasticsearch = client
	return nil
}

func (c *_elasticsearch) CreateIndex(ctx context.Context, indexes ...interfaces.ElasticsearchIndex) error {
	for i := 0; i < len(indexes); i++ {
		indexName := indexes[i].ElasticsearchIndexName()
		exists, err := global.ElasticsearchTypedClient.Indices.Exists(indexName).IsSuccess(ctx)
		if err != nil {
			return errors.Wrapf(err, "[elasticsearch][index:%s]索引查找失败!", indexName)
		}
		if !exists { // 索引不存在
			var response *create.Response
			response, err = global.ElasticsearchTypedClient.Indices.Create(indexName).Mappings(&types.TypeMapping{
				Properties: indexes[i].ElasticsearchProperties(),
			}).Do(ctx)
			if err != nil {
				return errors.Wrapf(err, "[elasticsearch][index:%s]创建失败!", indexName)
			}
			if !response.Acknowledged && response.Index != indexName {
				return errors.Wrapf(err, "[elasticsearch][index:%s]创建失败!", indexName)
			}
			fmt.Printf("[elasticsearch][index:%s]创建成功!\n", indexName)
		}
		if exists {
			fmt.Printf("[elasticsearch][index:%s]已存在!\n", indexName)
		}
	}
	return nil
}
