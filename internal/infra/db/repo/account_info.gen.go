// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repo

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/HarryWang29/echo_mind/internal/infra/db/model"
)

func newAccountInfo(db *gorm.DB, opts ...gen.DOOption) accountInfo {
	_accountInfo := accountInfo{}

	_accountInfo.accountInfoDo.UseDB(db, opts...)
	_accountInfo.accountInfoDo.UseModel(&model.AccountInfo{})

	tableName := _accountInfo.accountInfoDo.TableName()
	_accountInfo.ALL = field.NewAsterisk(tableName)
	_accountInfo.ID = field.NewInt64(tableName, "id")
	_accountInfo.UserName = field.NewString(tableName, "user_name")
	_accountInfo.Nickname = field.NewString(tableName, "nickname")
	_accountInfo.Hash = field.NewString(tableName, "hash")

	_accountInfo.fillFieldMap()

	return _accountInfo
}

type accountInfo struct {
	accountInfoDo accountInfoDo

	ALL      field.Asterisk
	ID       field.Int64
	UserName field.String
	Nickname field.String
	Hash     field.String

	fieldMap map[string]field.Expr
}

func (a accountInfo) Table(newTableName string) *accountInfo {
	a.accountInfoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountInfo) As(alias string) *accountInfo {
	a.accountInfoDo.DO = *(a.accountInfoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountInfo) updateTableName(table string) *accountInfo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UserName = field.NewString(table, "user_name")
	a.Nickname = field.NewString(table, "nickname")
	a.Hash = field.NewString(table, "hash")

	a.fillFieldMap()

	return a
}

func (a *accountInfo) WithContext(ctx context.Context) IAccountInfoDo {
	return a.accountInfoDo.WithContext(ctx)
}

func (a accountInfo) TableName() string { return a.accountInfoDo.TableName() }

func (a accountInfo) Alias() string { return a.accountInfoDo.Alias() }

func (a accountInfo) Columns(cols ...field.Expr) gen.Columns { return a.accountInfoDo.Columns(cols...) }

func (a *accountInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountInfo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_name"] = a.UserName
	a.fieldMap["nickname"] = a.Nickname
	a.fieldMap["hash"] = a.Hash
}

func (a accountInfo) clone(db *gorm.DB) accountInfo {
	a.accountInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountInfo) replaceDB(db *gorm.DB) accountInfo {
	a.accountInfoDo.ReplaceDB(db)
	return a
}

type accountInfoDo struct{ gen.DO }

type IAccountInfoDo interface {
	gen.SubQuery
	Debug() IAccountInfoDo
	WithContext(ctx context.Context) IAccountInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAccountInfoDo
	WriteDB() IAccountInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAccountInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAccountInfoDo
	Not(conds ...gen.Condition) IAccountInfoDo
	Or(conds ...gen.Condition) IAccountInfoDo
	Select(conds ...field.Expr) IAccountInfoDo
	Where(conds ...gen.Condition) IAccountInfoDo
	Order(conds ...field.Expr) IAccountInfoDo
	Distinct(cols ...field.Expr) IAccountInfoDo
	Omit(cols ...field.Expr) IAccountInfoDo
	Join(table schema.Tabler, on ...field.Expr) IAccountInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAccountInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAccountInfoDo
	Group(cols ...field.Expr) IAccountInfoDo
	Having(conds ...gen.Condition) IAccountInfoDo
	Limit(limit int) IAccountInfoDo
	Offset(offset int) IAccountInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountInfoDo
	Unscoped() IAccountInfoDo
	Create(values ...*model.AccountInfo) error
	CreateInBatches(values []*model.AccountInfo, batchSize int) error
	Save(values ...*model.AccountInfo) error
	First() (*model.AccountInfo, error)
	Take() (*model.AccountInfo, error)
	Last() (*model.AccountInfo, error)
	Find() ([]*model.AccountInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AccountInfo, err error)
	FindInBatches(result *[]*model.AccountInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AccountInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAccountInfoDo
	Assign(attrs ...field.AssignExpr) IAccountInfoDo
	Joins(fields ...field.RelationField) IAccountInfoDo
	Preload(fields ...field.RelationField) IAccountInfoDo
	FirstOrInit() (*model.AccountInfo, error)
	FirstOrCreate() (*model.AccountInfo, error)
	FindByPage(offset int, limit int) (result []*model.AccountInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAccountInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int) (result model.AccountInfo, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (a accountInfoDo) GetByID(id int) (result model.AccountInfo, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM account_info WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a accountInfoDo) Debug() IAccountInfoDo {
	return a.withDO(a.DO.Debug())
}

func (a accountInfoDo) WithContext(ctx context.Context) IAccountInfoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountInfoDo) ReadDB() IAccountInfoDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountInfoDo) WriteDB() IAccountInfoDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountInfoDo) Session(config *gorm.Session) IAccountInfoDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountInfoDo) Clauses(conds ...clause.Expression) IAccountInfoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountInfoDo) Returning(value interface{}, columns ...string) IAccountInfoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountInfoDo) Not(conds ...gen.Condition) IAccountInfoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountInfoDo) Or(conds ...gen.Condition) IAccountInfoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountInfoDo) Select(conds ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountInfoDo) Where(conds ...gen.Condition) IAccountInfoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountInfoDo) Order(conds ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountInfoDo) Distinct(cols ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountInfoDo) Omit(cols ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountInfoDo) Join(table schema.Tabler, on ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountInfoDo) Group(cols ...field.Expr) IAccountInfoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountInfoDo) Having(conds ...gen.Condition) IAccountInfoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountInfoDo) Limit(limit int) IAccountInfoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountInfoDo) Offset(offset int) IAccountInfoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountInfoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountInfoDo) Unscoped() IAccountInfoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountInfoDo) Create(values ...*model.AccountInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountInfoDo) CreateInBatches(values []*model.AccountInfo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountInfoDo) Save(values ...*model.AccountInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountInfoDo) First() (*model.AccountInfo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) Take() (*model.AccountInfo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) Last() (*model.AccountInfo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) Find() ([]*model.AccountInfo, error) {
	result, err := a.DO.Find()
	return result.([]*model.AccountInfo), err
}

func (a accountInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AccountInfo, err error) {
	buf := make([]*model.AccountInfo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountInfoDo) FindInBatches(result *[]*model.AccountInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountInfoDo) Attrs(attrs ...field.AssignExpr) IAccountInfoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountInfoDo) Assign(attrs ...field.AssignExpr) IAccountInfoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountInfoDo) Joins(fields ...field.RelationField) IAccountInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountInfoDo) Preload(fields ...field.RelationField) IAccountInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountInfoDo) FirstOrInit() (*model.AccountInfo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) FirstOrCreate() (*model.AccountInfo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountInfo), nil
	}
}

func (a accountInfoDo) FindByPage(offset int, limit int) (result []*model.AccountInfo, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a accountInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountInfoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountInfoDo) Delete(models ...*model.AccountInfo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountInfoDo) withDO(do gen.Dao) *accountInfoDo {
	a.DO = *do.(*gen.DO)
	return a
}
