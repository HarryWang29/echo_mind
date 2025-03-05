package group

import (
	"context"
	"fmt"
	"github.com/HarryWang29/echo_mind/config"
	"github.com/HarryWang29/echo_mind/internal/domain/account_info"
	"github.com/HarryWang29/echo_mind/internal/infra/db/model"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite/group"
	"github.com/HarryWang29/echo_mind/pkg/util"
	"gorm.io/gorm/clause"
	"path"
)

type Contact struct {
	groupContactDo repo.IGroupContactDo
	accountInfo    *account_info.AccountInfo
	wechat         *config.WechatConfig
	sqlInfo        []*sqliteInfo
}

type sqliteInfo struct {
	query                *group.Query
	groupContactDo       group.IGroupContactDo
	groupContactPersonDo group.IGroupMemberDo
	id                   string
	hash                 string
}

const dbName = "group_new.db"

func NewContact(w *config.WechatConfig, q *repo.Query, acc *account_info.AccountInfo) (c *Contact, err error) {
	c = &Contact{
		groupContactDo: q.GroupContact.WithContext(context.Background()),
		accountInfo:    acc,
		wechat:         w,
	}
	for _, info := range w.WatchInfo {
		db, err := sqlite.NewSQLite(w.Key, path.Join(info.Path, "Group"), dbName)
		if err != nil {
			return nil, err
		}
		query := group.Use(db.DB())
		c.sqlInfo = append(c.sqlInfo, &sqliteInfo{
			query:          query,
			groupContactDo: query.GroupContact.WithContext(context.Background()),
			id:             info.Id,
			hash:           info.Hash,
		})
	}
	return c, nil
}

func (c *Contact) Sync() error {
	for _, info := range c.sqlInfo {
		account, err := c.accountInfo.FindByHash(info.hash)
		if err != nil {
			return fmt.Errorf("find account(%s) info by hash: %w", info.hash, err)
		}
		err = c.SyncContact(account.ID, info)
		if err != nil {
			return fmt.Errorf("sync account(%s) contact person: %w", info.hash, err)
		}
	}
	return nil
}

func (c *Contact) SyncContact(accId int64, info *sqliteInfo) error {
	contacts, err := info.groupContactDo.Find()
	if err != nil {
		return fmt.Errorf("accountId(%d) groupContactDo.Find(): %w", accId, err)
	}
	todo := make([]*model.GroupContact, 0, len(contacts))
	for _, v := range contacts {
		cp := &model.GroupContact{
			AccountID:    accId,
			UserName:     v.MNsUsrName,
			Nickname:     v.Nickname,
			HeadImgURL:   v.MNsHeadImgURL,
			HeadHdImgURL: v.MNsHeadHDImgURL,
			GroupMember:  v.MNsChatRoomMemList,
			DbName:       dbName,
			Status:       false,
			Hash:         util.HashHex(util.MD5, v.MNsUsrName),
		}
		todo = append(todo, cp)
	}

	err = c.groupContactDo.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(todo, 2000)
	if err != nil {
		return fmt.Errorf("groupContactDo.CreateInBatches(): %w", err)
	}
	return nil
}
