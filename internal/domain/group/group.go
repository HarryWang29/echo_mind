package group

import (
	"context"
	"github.com/HarryWang29/echo_mind/config"
	"github.com/HarryWang29/echo_mind/internal/domain/account_info"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite/group"
	"path"
)

type Group struct {
	contactDo       repo.IGroupContactDo
	contactPersonDo repo.IGroupContactPersonDo
	userRelation    repo.IGroupUserRelationDo
	accountInfo     *account_info.AccountInfo
	wechat          *config.WechatWatchInfo
	sqlInfo         *sqliteInfo
}
type sqliteInfo struct {
	query           *group.Query
	contactDo       group.IGroupContactDo
	contactPersonDo group.IGroupMemberDo
	userRelation    group.IGroupUserRelationDo

	id   string
	hash string
}

const dbName = "group_new.db"

func New(w *config.WechatWatchInfo, q *repo.Query) (c *Group, err error) {
	c = &Group{
		contactDo:       q.GroupContact.WithContext(context.Background()),
		contactPersonDo: q.GroupContactPerson.WithContext(context.Background()),
		userRelation:    q.GroupUserRelation.WithContext(context.Background()),
		wechat:          w,
	}
	db, err := sqlite.NewSQLite(w.Key, path.Join(w.Path, "Group"), dbName)
	if err != nil {
		return nil, err
	}
	query := group.Use(db.DB())
	c.sqlInfo = &sqliteInfo{
		query:           query,
		contactDo:       query.GroupContact.WithContext(context.Background()),
		contactPersonDo: query.GroupMember.WithContext(context.Background()),
		userRelation:    query.GroupUserRelation.WithContext(context.Background()),
		id:              w.Id,
		hash:            w.Hash,
	}
	return c, nil
}
