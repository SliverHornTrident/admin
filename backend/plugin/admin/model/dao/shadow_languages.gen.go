// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"github.com/SliverHornTrident/shadow/plugin/admin/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newLanguage(db *gorm.DB, opts ...gen.DOOption) language {
	_language := language{}

	_language.languageDo.UseDB(db, opts...)
	_language.languageDo.UseModel(&model.Language{})

	tableName := _language.languageDo.TableName()
	_language.ALL = field.NewAsterisk(tableName)
	_language.Tag = field.NewString(tableName, "tag")
	_language.Path = field.NewString(tableName, "path")
	_language.Name = field.NewString(tableName, "name")
	_language.Unmarshal = field.NewString(tableName, "unmarshal")
	_language.Enable = field.NewBool(tableName, "enable")
	_language.Messages = languageHasManyMessages{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Messages", "model.LanguageMessage"),
	}

	_language.fillFieldMap()

	return _language
}

type language struct {
	languageDo

	ALL       field.Asterisk
	Tag       field.String
	Path      field.String
	Name      field.String
	Unmarshal field.String
	Enable    field.Bool
	Messages  languageHasManyMessages

	fieldMap map[string]field.Expr
}

func (l language) Table(newTableName string) *language {
	l.languageDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l language) As(alias string) *language {
	l.languageDo.DO = *(l.languageDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *language) updateTableName(table string) *language {
	l.ALL = field.NewAsterisk(table)
	l.Tag = field.NewString(table, "tag")
	l.Path = field.NewString(table, "path")
	l.Name = field.NewString(table, "name")
	l.Unmarshal = field.NewString(table, "unmarshal")
	l.Enable = field.NewBool(table, "enable")

	l.fillFieldMap()

	return l
}

func (l *language) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *language) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["tag"] = l.Tag
	l.fieldMap["path"] = l.Path
	l.fieldMap["name"] = l.Name
	l.fieldMap["unmarshal"] = l.Unmarshal
	l.fieldMap["enable"] = l.Enable

}

func (l language) clone(db *gorm.DB) language {
	l.languageDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l language) replaceDB(db *gorm.DB) language {
	l.languageDo.ReplaceDB(db)
	return l
}

type languageHasManyMessages struct {
	db *gorm.DB

	field.RelationField
}

func (a languageHasManyMessages) Where(conds ...field.Expr) *languageHasManyMessages {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a languageHasManyMessages) WithContext(ctx context.Context) *languageHasManyMessages {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a languageHasManyMessages) Session(session *gorm.Session) *languageHasManyMessages {
	a.db = a.db.Session(session)
	return &a
}

func (a languageHasManyMessages) Model(m *model.Language) *languageHasManyMessagesTx {
	return &languageHasManyMessagesTx{a.db.Model(m).Association(a.Name())}
}

type languageHasManyMessagesTx struct{ tx *gorm.Association }

func (a languageHasManyMessagesTx) Find() (result []*model.LanguageMessage, err error) {
	return result, a.tx.Find(&result)
}

func (a languageHasManyMessagesTx) Append(values ...*model.LanguageMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a languageHasManyMessagesTx) Replace(values ...*model.LanguageMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a languageHasManyMessagesTx) Delete(values ...*model.LanguageMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a languageHasManyMessagesTx) Clear() error {
	return a.tx.Clear()
}

func (a languageHasManyMessagesTx) Count() int64 {
	return a.tx.Count()
}

type languageDo struct{ gen.DO }

type ILanguageDo interface {
	gen.SubQuery
	Debug() ILanguageDo
	WithContext(ctx context.Context) ILanguageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILanguageDo
	WriteDB() ILanguageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILanguageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILanguageDo
	Not(conds ...gen.Condition) ILanguageDo
	Or(conds ...gen.Condition) ILanguageDo
	Select(conds ...field.Expr) ILanguageDo
	Where(conds ...gen.Condition) ILanguageDo
	Order(conds ...field.Expr) ILanguageDo
	Distinct(cols ...field.Expr) ILanguageDo
	Omit(cols ...field.Expr) ILanguageDo
	Join(table schema.Tabler, on ...field.Expr) ILanguageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILanguageDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILanguageDo
	Group(cols ...field.Expr) ILanguageDo
	Having(conds ...gen.Condition) ILanguageDo
	Limit(limit int) ILanguageDo
	Offset(offset int) ILanguageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILanguageDo
	Unscoped() ILanguageDo
	Create(values ...*model.Language) error
	CreateInBatches(values []*model.Language, batchSize int) error
	Save(values ...*model.Language) error
	First() (*model.Language, error)
	Take() (*model.Language, error)
	Last() (*model.Language, error)
	Find() ([]*model.Language, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Language, err error)
	FindInBatches(result *[]*model.Language, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Language) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILanguageDo
	Assign(attrs ...field.AssignExpr) ILanguageDo
	Joins(fields ...field.RelationField) ILanguageDo
	Preload(fields ...field.RelationField) ILanguageDo
	FirstOrInit() (*model.Language, error)
	FirstOrCreate() (*model.Language, error)
	FindByPage(offset int, limit int) (result []*model.Language, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILanguageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l languageDo) Debug() ILanguageDo {
	return l.withDO(l.DO.Debug())
}

func (l languageDo) WithContext(ctx context.Context) ILanguageDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l languageDo) ReadDB() ILanguageDo {
	return l.Clauses(dbresolver.Read)
}

func (l languageDo) WriteDB() ILanguageDo {
	return l.Clauses(dbresolver.Write)
}

func (l languageDo) Session(config *gorm.Session) ILanguageDo {
	return l.withDO(l.DO.Session(config))
}

func (l languageDo) Clauses(conds ...clause.Expression) ILanguageDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l languageDo) Returning(value interface{}, columns ...string) ILanguageDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l languageDo) Not(conds ...gen.Condition) ILanguageDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l languageDo) Or(conds ...gen.Condition) ILanguageDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l languageDo) Select(conds ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l languageDo) Where(conds ...gen.Condition) ILanguageDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l languageDo) Order(conds ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l languageDo) Distinct(cols ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l languageDo) Omit(cols ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l languageDo) Join(table schema.Tabler, on ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l languageDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l languageDo) RightJoin(table schema.Tabler, on ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l languageDo) Group(cols ...field.Expr) ILanguageDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l languageDo) Having(conds ...gen.Condition) ILanguageDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l languageDo) Limit(limit int) ILanguageDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l languageDo) Offset(offset int) ILanguageDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l languageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILanguageDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l languageDo) Unscoped() ILanguageDo {
	return l.withDO(l.DO.Unscoped())
}

func (l languageDo) Create(values ...*model.Language) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l languageDo) CreateInBatches(values []*model.Language, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l languageDo) Save(values ...*model.Language) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l languageDo) First() (*model.Language, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Language), nil
	}
}

func (l languageDo) Take() (*model.Language, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Language), nil
	}
}

func (l languageDo) Last() (*model.Language, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Language), nil
	}
}

func (l languageDo) Find() ([]*model.Language, error) {
	result, err := l.DO.Find()
	return result.([]*model.Language), err
}

func (l languageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Language, err error) {
	buf := make([]*model.Language, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l languageDo) FindInBatches(result *[]*model.Language, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l languageDo) Attrs(attrs ...field.AssignExpr) ILanguageDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l languageDo) Assign(attrs ...field.AssignExpr) ILanguageDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l languageDo) Joins(fields ...field.RelationField) ILanguageDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l languageDo) Preload(fields ...field.RelationField) ILanguageDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l languageDo) FirstOrInit() (*model.Language, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Language), nil
	}
}

func (l languageDo) FirstOrCreate() (*model.Language, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Language), nil
	}
}

func (l languageDo) FindByPage(offset int, limit int) (result []*model.Language, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l languageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l languageDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l languageDo) Delete(models ...*model.Language) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *languageDo) withDO(do gen.Dao) *languageDo {
	l.DO = *do.(*gen.DO)
	return l
}
