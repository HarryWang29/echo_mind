// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package group

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                 = new(Query)
	GroupContact      *groupContact
	GroupMember       *groupMember
	GroupUserRelation *groupUserRelation
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	GroupContact = &Q.GroupContact
	GroupMember = &Q.GroupMember
	GroupUserRelation = &Q.GroupUserRelation
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                db,
		GroupContact:      newGroupContact(db, opts...),
		GroupMember:       newGroupMember(db, opts...),
		GroupUserRelation: newGroupUserRelation(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	GroupContact      groupContact
	GroupMember       groupMember
	GroupUserRelation groupUserRelation
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		GroupContact:      q.GroupContact.clone(db),
		GroupMember:       q.GroupMember.clone(db),
		GroupUserRelation: q.GroupUserRelation.clone(db),
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
		db:                db,
		GroupContact:      q.GroupContact.replaceDB(db),
		GroupMember:       q.GroupMember.replaceDB(db),
		GroupUserRelation: q.GroupUserRelation.replaceDB(db),
	}
}

type queryCtx struct {
	GroupContact      IGroupContactDo
	GroupMember       IGroupMemberDo
	GroupUserRelation IGroupUserRelationDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		GroupContact:      q.GroupContact.WithContext(ctx),
		GroupMember:       q.GroupMember.WithContext(ctx),
		GroupUserRelation: q.GroupUserRelation.WithContext(ctx),
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
