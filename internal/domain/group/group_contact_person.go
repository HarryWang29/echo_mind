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

type ContactPerson struct {
	groupContactPersonDo repo.IGroupContactPersonDo
	accountInfo          *account_info.AccountInfo
	wechat               *config.WechatConfig
	sqlInfo              []*sqliteInfo
}

func NewContactPerson(w *config.WechatConfig, q *repo.Query, acc *account_info.AccountInfo) (c *Contact, err error) {
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
			query:                query,
			groupContactPersonDo: query.GroupMember.WithContext(context.Background()),
			id:                   info.Id,
			hash:                 info.Hash,
		})
	}
	return c, nil
}

func (c *ContactPerson) Sync() error {
	for _, info := range c.sqlInfo {
		account, err := c.accountInfo.FindByHash(info.hash)
		if err != nil {
			return fmt.Errorf("find account(%s) info by hash: %w", info.hash, err)
		}
		err = c.SyncContactPerson(account.ID, info)
		if err != nil {
			return fmt.Errorf("sync account(%s) contact person: %w", info.hash, err)
		}
	}
	return nil
}

func (c *ContactPerson) SyncContactPerson(accId int64, info *sqliteInfo) error {
	members, err := info.groupContactPersonDo.Find()
	if err != nil {
		return fmt.Errorf("accountId(%d) groupContactPersonDo.Find(): %w", accId, err)
	}
	todo := make([]*model.GroupContactPerson, 0, len(members))
	for _, v := range members {
		cp := &model.GroupContactPerson{
			AccountID:    accId,
			UserName:     v.MNsUsrName,
			Nickname:     v.Nickname,
			Remark:       v.MNsRemark,
			HeadImgURL:   v.MNsHeadImgURL,
			HeadHdImgURL: v.MNsHeadHDImgURL,
			Sex:          v.MUISex,
			Type:         v.MUIType,
			DbName:       dbName,
			Status:       false,
			Hash:         util.HashHex(util.MD5, v.MNsUsrName),
		}
		todo = append(todo, cp)
	}

	err = c.groupContactPersonDo.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(todo, 2000)
	if err != nil {
		return fmt.Errorf("groupContactPersonDo.CreateInBatches(): %w", err)
	}
	return nil
}
