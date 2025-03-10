// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repo

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	AccountInfo          *accountInfo
	ContactPerson        *contactPerson
	GroupContact         *groupContact
	GroupContactPerson   *groupContactPerson
	GroupUserRelation    *groupUserRelation
	Message              *message
	SessionAbstract      *sessionAbstract
	SessionAbstractBrand *sessionAbstractBrand
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AccountInfo = &Q.AccountInfo
	ContactPerson = &Q.ContactPerson
	GroupContact = &Q.GroupContact
	GroupContactPerson = &Q.GroupContactPerson
	GroupUserRelation = &Q.GroupUserRelation
	Message = &Q.Message
	SessionAbstract = &Q.SessionAbstract
	SessionAbstractBrand = &Q.SessionAbstractBrand
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		AccountInfo:          newAccountInfo(db, opts...),
		ContactPerson:        newContactPerson(db, opts...),
		GroupContact:         newGroupContact(db, opts...),
		GroupContactPerson:   newGroupContactPerson(db, opts...),
		GroupUserRelation:    newGroupUserRelation(db, opts...),
		Message:              newMessage(db, opts...),
		SessionAbstract:      newSessionAbstract(db, opts...),
		SessionAbstractBrand: newSessionAbstractBrand(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AccountInfo          accountInfo
	ContactPerson        contactPerson
	GroupContact         groupContact
	GroupContactPerson   groupContactPerson
	GroupUserRelation    groupUserRelation
	Message              message
	SessionAbstract      sessionAbstract
	SessionAbstractBrand sessionAbstractBrand
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		AccountInfo:          q.AccountInfo.clone(db),
		ContactPerson:        q.ContactPerson.clone(db),
		GroupContact:         q.GroupContact.clone(db),
		GroupContactPerson:   q.GroupContactPerson.clone(db),
		GroupUserRelation:    q.GroupUserRelation.clone(db),
		Message:              q.Message.clone(db),
		SessionAbstract:      q.SessionAbstract.clone(db),
		SessionAbstractBrand: q.SessionAbstractBrand.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		AccountInfo:          q.AccountInfo.replaceDB(db),
		ContactPerson:        q.ContactPerson.replaceDB(db),
		GroupContact:         q.GroupContact.replaceDB(db),
		GroupContactPerson:   q.GroupContactPerson.replaceDB(db),
		GroupUserRelation:    q.GroupUserRelation.replaceDB(db),
		Message:              q.Message.replaceDB(db),
		SessionAbstract:      q.SessionAbstract.replaceDB(db),
		SessionAbstractBrand: q.SessionAbstractBrand.replaceDB(db),
	}
}

type queryCtx struct {
	AccountInfo          IAccountInfoDo
	ContactPerson        IContactPersonDo
	GroupContact         IGroupContactDo
	GroupContactPerson   IGroupContactPersonDo
	GroupUserRelation    IGroupUserRelationDo
	Message              IMessageDo
	SessionAbstract      ISessionAbstractDo
	SessionAbstractBrand ISessionAbstractBrandDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AccountInfo:          q.AccountInfo.WithContext(ctx),
		ContactPerson:        q.ContactPerson.WithContext(ctx),
		GroupContact:         q.GroupContact.WithContext(ctx),
		GroupContactPerson:   q.GroupContactPerson.WithContext(ctx),
		GroupUserRelation:    q.GroupUserRelation.WithContext(ctx),
		Message:              q.Message.WithContext(ctx),
		SessionAbstract:      q.SessionAbstract.WithContext(ctx),
		SessionAbstractBrand: q.SessionAbstractBrand.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
