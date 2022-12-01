// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"

	pb "github.com/beihai0xff/pudding/api/gen/pudding/trigger/v1"
	"github.com/beihai0xff/pudding/app/trigger/repo/storage/po"
)

func newWebhookTriggerTemplate(db *gorm.DB, opts ...gen.DOOption) webhookTriggerTemplate {
	_webhookTriggerTemplate := webhookTriggerTemplate{}

	_webhookTriggerTemplate.webhookTriggerTemplateDo.UseDB(db, opts...)
	_webhookTriggerTemplate.webhookTriggerTemplateDo.UseModel(&po.WebhookTriggerTemplate{})

	tableName := _webhookTriggerTemplate.webhookTriggerTemplateDo.TableName()
	_webhookTriggerTemplate.ALL = field.NewAsterisk(tableName)
	_webhookTriggerTemplate.ID = field.NewUint(tableName, "id")
	_webhookTriggerTemplate.CreatedAt = field.NewTime(tableName, "created_at")
	_webhookTriggerTemplate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_webhookTriggerTemplate.DeletedAt = field.NewField(tableName, "deleted_at")
	_webhookTriggerTemplate.Topic = field.NewString(tableName, "topic")
	_webhookTriggerTemplate.Payload = field.NewBytes(tableName, "payload")
	_webhookTriggerTemplate.ExceptedEndTime = field.NewTime(tableName, "excepted_end_time")
	_webhookTriggerTemplate.ExceptedLoopTimes = field.NewUint64(tableName, "excepted_loop_times")
	_webhookTriggerTemplate.LoopedTimes = field.NewUint64(tableName, "looped_times")
	_webhookTriggerTemplate.Status = field.NewInt32(tableName, "status")

	_webhookTriggerTemplate.fillFieldMap()

	return _webhookTriggerTemplate
}

type webhookTriggerTemplate struct {
	webhookTriggerTemplateDo webhookTriggerTemplateDo

	ALL               field.Asterisk
	ID                field.Uint
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	Topic             field.String
	Payload           field.Bytes
	ExceptedEndTime   field.Time
	ExceptedLoopTimes field.Uint64
	LoopedTimes       field.Uint64
	Status            field.Int32

	fieldMap map[string]field.Expr
}

func (w webhookTriggerTemplate) Table(newTableName string) *webhookTriggerTemplate {
	w.webhookTriggerTemplateDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w webhookTriggerTemplate) As(alias string) *webhookTriggerTemplate {
	w.webhookTriggerTemplateDo.DO = *(w.webhookTriggerTemplateDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *webhookTriggerTemplate) updateTableName(table string) *webhookTriggerTemplate {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.Topic = field.NewString(table, "topic")
	w.Payload = field.NewBytes(table, "payload")
	w.ExceptedEndTime = field.NewTime(table, "excepted_end_time")
	w.ExceptedLoopTimes = field.NewUint64(table, "excepted_loop_times")
	w.LoopedTimes = field.NewUint64(table, "looped_times")
	w.Status = field.NewInt32(table, "status")

	w.fillFieldMap()

	return w
}

func (w *webhookTriggerTemplate) WithContext(ctx context.Context) *webhookTriggerTemplateDo {
	return w.webhookTriggerTemplateDo.WithContext(ctx)
}

func (w webhookTriggerTemplate) TableName() string { return w.webhookTriggerTemplateDo.TableName() }

func (w webhookTriggerTemplate) Alias() string { return w.webhookTriggerTemplateDo.Alias() }

func (w *webhookTriggerTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *webhookTriggerTemplate) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 10)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["topic"] = w.Topic
	w.fieldMap["payload"] = w.Payload
	w.fieldMap["excepted_end_time"] = w.ExceptedEndTime
	w.fieldMap["excepted_loop_times"] = w.ExceptedLoopTimes
	w.fieldMap["looped_times"] = w.LoopedTimes
	w.fieldMap["status"] = w.Status
}

func (w webhookTriggerTemplate) clone(db *gorm.DB) webhookTriggerTemplate {
	w.webhookTriggerTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w webhookTriggerTemplate) replaceDB(db *gorm.DB) webhookTriggerTemplate {
	w.webhookTriggerTemplateDo.ReplaceDB(db)
	return w
}

type webhookTriggerTemplateDo struct{ gen.DO }

// SELECT * FROM @@table WHERE id=@id
func (w webhookTriggerTemplateDo) FindByID(id uint) (result *po.WebhookTriggerTemplate, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM webhook_trigger_template WHERE id=? ")

	var executeSQL *gorm.DB

	executeSQL = w.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result)
	err = executeSQL.Error
	return
}

// UPDATE @@table
// {{set}}
//
//	{{if status > 0}} status=@status, {{end}}
//
// {{end}}
// WHERE id=@id
func (w webhookTriggerTemplateDo) UpdateStatus(ctx context.Context, id uint, status pb.TriggerStatus) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE webhook_trigger_template ")
	var setSQL0 strings.Builder
	if status > 0 {
		params = append(params, status)
		setSQL0.WriteString("status=?, ")
	}
	helper.JoinSetBuilder(&generateSQL, setSQL0)
	params = append(params, id)
	generateSQL.WriteString("WHERE id=? ")

	var executeSQL *gorm.DB

	executeSQL = w.UnderlyingDB().Exec(generateSQL.String(), params...)
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error
	return
}

func (w webhookTriggerTemplateDo) Debug() *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Debug())
}

func (w webhookTriggerTemplateDo) WithContext(ctx context.Context) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w webhookTriggerTemplateDo) ReadDB() *webhookTriggerTemplateDo {
	return w.Clauses(dbresolver.Read)
}

func (w webhookTriggerTemplateDo) WriteDB() *webhookTriggerTemplateDo {
	return w.Clauses(dbresolver.Write)
}

func (w webhookTriggerTemplateDo) Session(config *gorm.Session) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Session(config))
}

func (w webhookTriggerTemplateDo) Clauses(conds ...clause.Expression) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w webhookTriggerTemplateDo) Returning(value interface{}, columns ...string) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w webhookTriggerTemplateDo) Not(conds ...gen.Condition) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w webhookTriggerTemplateDo) Or(conds ...gen.Condition) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w webhookTriggerTemplateDo) Select(conds ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w webhookTriggerTemplateDo) Where(conds ...gen.Condition) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w webhookTriggerTemplateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *webhookTriggerTemplateDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w webhookTriggerTemplateDo) Order(conds ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w webhookTriggerTemplateDo) Distinct(cols ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w webhookTriggerTemplateDo) Omit(cols ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w webhookTriggerTemplateDo) Join(table schema.Tabler, on ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w webhookTriggerTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w webhookTriggerTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w webhookTriggerTemplateDo) Group(cols ...field.Expr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w webhookTriggerTemplateDo) Having(conds ...gen.Condition) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w webhookTriggerTemplateDo) Limit(limit int) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w webhookTriggerTemplateDo) Offset(offset int) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w webhookTriggerTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w webhookTriggerTemplateDo) Unscoped() *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Unscoped())
}

func (w webhookTriggerTemplateDo) Create(values ...*po.WebhookTriggerTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w webhookTriggerTemplateDo) CreateInBatches(values []*po.WebhookTriggerTemplate, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w webhookTriggerTemplateDo) Save(values ...*po.WebhookTriggerTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w webhookTriggerTemplateDo) First() (*po.WebhookTriggerTemplate, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.WebhookTriggerTemplate), nil
	}
}

func (w webhookTriggerTemplateDo) Take() (*po.WebhookTriggerTemplate, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.WebhookTriggerTemplate), nil
	}
}

func (w webhookTriggerTemplateDo) Last() (*po.WebhookTriggerTemplate, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.WebhookTriggerTemplate), nil
	}
}

func (w webhookTriggerTemplateDo) Find() ([]*po.WebhookTriggerTemplate, error) {
	result, err := w.DO.Find()
	return result.([]*po.WebhookTriggerTemplate), err
}

func (w webhookTriggerTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.WebhookTriggerTemplate, err error) {
	buf := make([]*po.WebhookTriggerTemplate, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w webhookTriggerTemplateDo) FindInBatches(result *[]*po.WebhookTriggerTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w webhookTriggerTemplateDo) Attrs(attrs ...field.AssignExpr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w webhookTriggerTemplateDo) Assign(attrs ...field.AssignExpr) *webhookTriggerTemplateDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w webhookTriggerTemplateDo) Joins(fields ...field.RelationField) *webhookTriggerTemplateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w webhookTriggerTemplateDo) Preload(fields ...field.RelationField) *webhookTriggerTemplateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w webhookTriggerTemplateDo) FirstOrInit() (*po.WebhookTriggerTemplate, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.WebhookTriggerTemplate), nil
	}
}

func (w webhookTriggerTemplateDo) FirstOrCreate() (*po.WebhookTriggerTemplate, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.WebhookTriggerTemplate), nil
	}
}

func (w webhookTriggerTemplateDo) FindByPage(offset int, limit int) (result []*po.WebhookTriggerTemplate, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w webhookTriggerTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w webhookTriggerTemplateDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w webhookTriggerTemplateDo) Delete(models ...*po.WebhookTriggerTemplate) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *webhookTriggerTemplateDo) withDO(do gen.Dao) *webhookTriggerTemplateDo {
	w.DO = *do.(*gen.DO)
	return w
}
