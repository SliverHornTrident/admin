//go:build (gorm || gorms) && (tidb || mysql || postgres || sqlite || clickhouse || mssql || sqlserver || oracle)

package global

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
)

// Operator 操作者
type Operator struct {
	Created uint `json:"created" gorm:"column:created;comment:创建者"`
	Updated uint `json:"updated" gorm:"column:updated;comment:更新者"`
	Deleted uint `json:"deleted" gorm:"column:deleted;comment:删除者"`
}

func (o *Operator) BeforeCreate(tx *gorm.DB) error {
	userId := tx.Statement.Context.Value("Operator")
	value, ok := userId.(uint)
	if !ok {
		return errors.New("操作者Id类型错误!")
	}
	o.Created = value
	o.Updated = value
	return nil
}

func (o *Operator) BeforeUpdate(tx *gorm.DB) error {
	userId := tx.Statement.Context.Value("Operator")
	value, o1 := userId.(uint)
	if !o1 {
		return errors.New("操作者Id类型错误!")
	}
	where, o2 := tx.Statement.Clauses["WHERE"].Expression.(clause.Where)
	if !o2 {
		if tx.Statement.Model != nil {
			_, queryValues := schema.GetIdentityFieldValuesMap(tx.Statement.Context, reflect.ValueOf(tx.Statement.Model), tx.Statement.Schema.PrimaryFields)
			column, values := schema.ToQueryValues(tx.Statement.Table, tx.Statement.Schema.PrimaryFieldDBNames, queryValues)
			if len(values) > 0 {
				tx.Statement.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
			}
		}
		where, o2 = tx.Statement.Clauses["WHERE"].Expression.(clause.Where)
		if !o2 {
			zap.L().Error("删除时获取where条件失败!")
			return nil
		}
	}
	length := len(where.Exprs)
	query := make([]string, 0, length)
	args := make([]any, 0, length)
	for i := 0; i < length; i++ {
		eq, o3 := where.Exprs[i].(clause.Eq)
		if o3 {
			column, o4 := eq.Column.(clause.Column)
			if o4 {
				query = append(query, column.Name+" = ?")
				args = append(args, eq.Value)
			}
		}
		in, o5 := where.Exprs[i].(clause.IN)
		if o5 {
			column, o6 := in.Column.(clause.Column)
			if o6 {
				query = append(query, column.Name+" IN ?")
				args = append(args, in.Values)
			}
		}
	}
	if (len(query) != 0 && len(args) != 0) && len(query) == len(args) {
		err := tx.Table(tx.Statement.Table).Where(strings.Join(query, " AND "), args...).UpdateColumn("updated", value).Error
		if err != nil {
			return errors.New("更改更新者id失败!")
		}
	}
	return nil
}
func (o *Operator) BeforeDelete(tx *gorm.DB) error {
	userId := tx.Statement.Context.Value("Operator")
	value, o1 := userId.(uint)
	if !o1 {
		return errors.New("操作者Id类型错误!")
	}
	where, o2 := tx.Statement.Clauses["WHERE"].Expression.(clause.Where)
	if !o2 {
		if tx.Statement.Model != nil {
			_, queryValues := schema.GetIdentityFieldValuesMap(tx.Statement.Context, reflect.ValueOf(tx.Statement.Model), tx.Statement.Schema.PrimaryFields)
			column, values := schema.ToQueryValues(tx.Statement.Table, tx.Statement.Schema.PrimaryFieldDBNames, queryValues)
			if len(values) > 0 {
				tx.Statement.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
			}
		}
		where, o2 = tx.Statement.Clauses["WHERE"].Expression.(clause.Where)
		if !o2 {
			zap.L().Error("删除时获取where条件失败!")
			return nil
		}
	}
	length := len(where.Exprs)
	query := make([]string, 0, length)
	args := make([]any, 0, length)
	for i := 0; i < length; i++ {
		eq, o3 := where.Exprs[i].(clause.Eq)
		if o3 {
			column, o4 := eq.Column.(clause.Column)
			if o4 {
				query = append(query, column.Name+" = ?")
				args = append(args, eq.Value)
			}
			continue
		}
		in, o5 := where.Exprs[i].(clause.IN)
		if o5 {
			column, o6 := in.Column.(clause.Column)
			if o6 {
				query = append(query, column.Name+" IN ?")
				args = append(args, in.Values)
			}
		}
	}
	if (len(query) != 0 && len(args) != 0) && len(query) == len(args) {
		err := tx.Table(tx.Statement.Table).Where(strings.Join(query, " AND "), args...).UpdateColumn("deleted", value).Error
		if err != nil {
			return errors.New("更新删除者id失败!")
		}
	}
	return nil
}
