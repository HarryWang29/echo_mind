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
	groupContactDo       repo.IGroupContactDo
	groupContactPersonDo repo.IGroupContactPersonDo
	accountInfo          *account_info.AccountInfo
	wechat               *config.WechatConfig
	sqlInfo              []*sqliteInfo
}
type sqliteInfo struct {
	query                *group.Query
	groupContactDo       group.IGroupContactDo
	groupContactPersonDo group.IGroupMemberDo
	id                   string
	hash                 string
}

const dbName = "group_new.db"

func New(w *config.WechatConfig, q *repo.Query, acc *account_info.AccountInfo) (c *Group, err error) {
	c = &Group{
		groupContactDo:       q.GroupContact.WithContext(context.Background()),
		groupContactPersonDo: q.GroupContactPerson.WithContext(context.Background()),
		accountInfo:          acc,
		wechat:               w,
	}
	for _, info := range w.WatchInfo {
		db, err := sqlite.NewSQLite(w.Key, path.Join(info.Path, "Group"), dbName)
		if err != nil {
			return nil, err
		}
		query := group.Use(db.DB())
		c.sqlInfo = append(c.sqlInfo, &sqliteInfo{
			query:                query,
			groupContactDo:       query.GroupContact.WithContext(context.Background()),
			groupContactPersonDo: query.GroupMember.WithContext(context.Background()),
			id:                   info.Id,
			hash:                 info.Hash,
		})
	}
	return c, nil
}
