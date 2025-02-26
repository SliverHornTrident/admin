//go:build gorm || gorms

package constant

type GormsType string

const (
	GormsTypeTidb       GormsType = "tidb"
	GormsTypeMysql      GormsType = "mysql"
	GormsTypeSqlite     GormsType = "sqlite"
	GormsTypePostgres   GormsType = "postgres"
	GormsTypeClickhouse GormsType = "clickhouse"
)
