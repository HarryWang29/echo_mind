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

func newMessage(db *gorm.DB, opts ...gen.DOOption) message {
	_message := message{}

	_message.messageDo.UseDB(db, opts...)
	_message.messageDo.UseModel(&model.Message{})

	tableName := _message.messageDo.TableName()
	_message.ALL = field.NewAsterisk(tableName)
	_message.ID = field.NewInt64(tableName, "id")
	_message.AccountID = field.NewInt64(tableName, "account_id")
	_message.Hash = field.NewString(tableName, "hash")
	_message.LocalID = field.NewInt64(tableName, "local_id")
	_message.SvrID = field.NewInt64(tableName, "svr_id")
	_message.CreateTime = field.NewInt64(tableName, "create_time")
	_message.Content = field.NewString(tableName, "content")
	_message.Translate = field.NewString(tableName, "translate")
	_message.Status = field.NewInt64(tableName, "status")
	_message.ImgStatus = field.NewInt64(tableName, "img_status")
	_message.MessageType = field.NewInt32(tableName, "message_type")
	_message.Des = field.NewBool(tableName, "des")
	_message.Source = field.NewString(tableName, "source")
	_message.VoiceText = field.NewString(tableName, "voice_text")
	_message.Seq = field.NewInt64(tableName, "seq")
	_message.DbName = field.NewString(tableName, "db_name")

	_message.fillFieldMap()

	return _message
}

type message struct {
	messageDo messageDo

	ALL         field.Asterisk
	ID          field.Int64
	AccountID   field.Int64 // 账号id
	Hash        field.String
	LocalID     field.Int64
	SvrID       field.Int64  // 服务器id
	CreateTime  field.Int64  // 消息时间
	Content     field.String // 正文
	Translate   field.String
	Status      field.Int64 // 状态
	ImgStatus   field.Int64 // 图片状态
	MessageType field.Int32 // 类型
	Des         field.Bool
	Source      field.String
	VoiceText   field.String
	Seq         field.Int64
	DbName      field.String // 数据库名称

	fieldMap map[string]field.Expr
}

func (m message) Table(newTableName string) *message {
	m.messageDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m message) As(alias string) *message {
	m.messageDo.DO = *(m.messageDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *message) updateTableName(table string) *message {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.AccountID = field.NewInt64(table, "account_id")
	m.Hash = field.NewString(table, "hash")
	m.LocalID = field.NewInt64(table, "local_id")
	m.SvrID = field.NewInt64(table, "svr_id")
	m.CreateTime = field.NewInt64(table, "create_time")
	m.Content = field.NewString(table, "content")
	m.Translate = field.NewString(table, "translate")
	m.Status = field.NewInt64(table, "status")
	m.ImgStatus = field.NewInt64(table, "img_status")
	m.MessageType = field.NewInt32(table, "message_type")
	m.Des = field.NewBool(table, "des")
	m.Source = field.NewString(table, "source")
	m.VoiceText = field.NewString(table, "voice_text")
	m.Seq = field.NewInt64(table, "seq")
	m.DbName = field.NewString(table, "db_name")

	m.fillFieldMap()

	return m
}

func (m *message) WithContext(ctx context.Context) IMessageDo { return m.messageDo.WithContext(ctx) }

func (m message) TableName() string { return m.messageDo.TableName() }

func (m message) Alias() string { return m.messageDo.Alias() }

func (m message) Columns(cols ...field.Expr) gen.Columns { return m.messageDo.Columns(cols...) }

func (m *message) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *message) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 16)
	m.fieldMap["id"] = m.ID
	m.fieldMap["account_id"] = m.AccountID
	m.fieldMap["hash"] = m.Hash
	m.fieldMap["local_id"] = m.LocalID
	m.fieldMap["svr_id"] = m.SvrID
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["content"] = m.Content
	m.fieldMap["translate"] = m.Translate
	m.fieldMap["status"] = m.Status
	m.fieldMap["img_status"] = m.ImgStatus
	m.fieldMap["message_type"] = m.MessageType
	m.fieldMap["des"] = m.Des
	m.fieldMap["source"] = m.Source
	m.fieldMap["voice_text"] = m.VoiceText
	m.fieldMap["seq"] = m.Seq
	m.fieldMap["db_name"] = m.DbName
}

func (m message) clone(db *gorm.DB) message {
	m.messageDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m message) replaceDB(db *gorm.DB) message {
	m.messageDo.ReplaceDB(db)
	return m
}

type messageDo struct{ gen.DO }

type IMessageDo interface {
	gen.SubQuery
	Debug() IMessageDo
	WithContext(ctx context.Context) IMessageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMessageDo
	WriteDB() IMessageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMessageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMessageDo
	Not(conds ...gen.Condition) IMessageDo
	Or(conds ...gen.Condition) IMessageDo
	Select(conds ...field.Expr) IMessageDo
	Where(conds ...gen.Condition) IMessageDo
	Order(conds ...field.Expr) IMessageDo
	Distinct(cols ...field.Expr) IMessageDo
	Omit(cols ...field.Expr) IMessageDo
	Join(table schema.Tabler, on ...field.Expr) IMessageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMessageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMessageDo
	Group(cols ...field.Expr) IMessageDo
	Having(conds ...gen.Condition) IMessageDo
	Limit(limit int) IMessageDo
	Offset(offset int) IMessageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageDo
	Unscoped() IMessageDo
	Create(values ...*model.Message) error
	CreateInBatches(values []*model.Message, batchSize int) error
	Save(values ...*model.Message) error
	First() (*model.Message, error)
	Take() (*model.Message, error)
	Last() (*model.Message, error)
	Find() ([]*model.Message, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Message, err error)
	FindInBatches(result *[]*model.Message, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Message) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMessageDo
	Assign(attrs ...field.AssignExpr) IMessageDo
	Joins(fields ...field.RelationField) IMessageDo
	Preload(fields ...field.RelationField) IMessageDo
	FirstOrInit() (*model.Message, error)
	FirstOrCreate() (*model.Message, error)
	FindByPage(offset int, limit int) (result []*model.Message, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMessageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int) (result model.Message, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (m messageDo) GetByID(id int) (result model.Message, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM message WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = m.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (m messageDo) Debug() IMessageDo {
	return m.withDO(m.DO.Debug())
}

func (m messageDo) WithContext(ctx context.Context) IMessageDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messageDo) ReadDB() IMessageDo {
	return m.Clauses(dbresolver.Read)
}

func (m messageDo) WriteDB() IMessageDo {
	return m.Clauses(dbresolver.Write)
}

func (m messageDo) Session(config *gorm.Session) IMessageDo {
	return m.withDO(m.DO.Session(config))
}

func (m messageDo) Clauses(conds ...clause.Expression) IMessageDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messageDo) Returning(value interface{}, columns ...string) IMessageDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messageDo) Not(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messageDo) Or(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messageDo) Select(conds ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messageDo) Where(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messageDo) Order(conds ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messageDo) Distinct(cols ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messageDo) Omit(cols ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messageDo) Join(table schema.Tabler, on ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMessageDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messageDo) RightJoin(table schema.Tabler, on ...field.Expr) IMessageDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messageDo) Group(cols ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messageDo) Having(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messageDo) Limit(limit int) IMessageDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messageDo) Offset(offset int) IMessageDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messageDo) Unscoped() IMessageDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messageDo) Create(values ...*model.Message) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messageDo) CreateInBatches(values []*model.Message, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messageDo) Save(values ...*model.Message) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messageDo) First() (*model.Message, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Message), nil
	}
}

func (m messageDo) Take() (*model.Message, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Message), nil
	}
}

func (m messageDo) Last() (*model.Message, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Message), nil
	}
}

func (m messageDo) Find() ([]*model.Message, error) {
	result, err := m.DO.Find()
	return result.([]*model.Message), err
}

func (m messageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Message, err error) {
	buf := make([]*model.Message, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messageDo) FindInBatches(result *[]*model.Message, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messageDo) Attrs(attrs ...field.AssignExpr) IMessageDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messageDo) Assign(attrs ...field.AssignExpr) IMessageDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messageDo) Joins(fields ...field.RelationField) IMessageDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messageDo) Preload(fields ...field.RelationField) IMessageDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messageDo) FirstOrInit() (*model.Message, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Message), nil
	}
}

func (m messageDo) FirstOrCreate() (*model.Message, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Message), nil
	}
}

func (m messageDo) FindByPage(offset int, limit int) (result []*model.Message, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m messageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messageDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messageDo) Delete(models ...*model.Message) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messageDo) withDO(do gen.Dao) *messageDo {
	m.DO = *do.(*gen.DO)
	return m
}
