//go:build elasticsearch

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	Elasticsearch            *elasticsearch.Client
	ElasticsearchConfig      config.Elasticsearch
	ElasticsearchTypedClient *elasticsearch.TypedClient
)
