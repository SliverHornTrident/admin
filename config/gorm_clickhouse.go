//go:build (gorm || gorms) && clickhouse

package config

import (
	"fmt"
	"gorm.io/driver/clickhouse"
)

func (c Gorm) Clickhouse() clickhouse.Config {
	if c.Config == "" {
		c.Config = "read_timeout=10&write_timeout=20"
	}
	return clickhouse.Config{
		DSN:                          fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&%s", c.Host, c.Port, c.Dbname, c.Username, c.Password, c.Config),
		DisableDatetimePrecision:     true,     // disable datetime64 precision, not supported before clickhouse 20.4
		DontSupportRenameColumn:      true,     // rename column not supported before clickhouse 20.4
		DontSupportEmptyDefaultValue: false,    // do not consider empty strings as valid default values
		SkipInitializeWithVersion:    false,    // smart configure based on used version
		DefaultGranularity:           3,        // 1 granule = 8192 rows
		DefaultCompression:           "LZ4",    // default compression algorithm. LZ4 is lossless
		DefaultIndexType:             "minmax", // index stores extremes of the expression
		DefaultTableEngineOpts:       "ENGINE=MergeTree() ORDER BY tuple()",
	}
}
