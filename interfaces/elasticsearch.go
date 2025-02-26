package interfaces

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type ElasticsearchIndex interface {
	ElasticsearchIndexName() string
	ElasticsearchProperties() map[string]types.Property
}
