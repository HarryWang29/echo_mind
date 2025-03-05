package contact_person

import (
	"context"
	"fmt"
	"github.com/HarryWang29/echo_mind/config"
	"github.com/HarryWang29/echo_mind/internal/domain/account_info"
	"github.com/HarryWang29/echo_mind/internal/infra/db/model"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite/contact"
	"github.com/HarryWang29/echo_mind/pkg/util"
	"gorm.io/gorm/clause"
	"path"
)

type ContactPerson struct {
	contactPersonDo repo.IContactPersonDo
	accountInfo     *account_info.AccountInfo
	wechat          *config.WechatConfig
	sqlInfo         []*sqliteInfo
}

type sqliteInfo struct {
	query       *contact.Query
	wcContactDo contact.IWCContactDo
	id          string
	hash        string
}

const dbName = "wccontact_new2.db"

func NewContactPerson(w *config.WechatConfig, q *repo.Query, acc *account_info.AccountInfo) (c *ContactPerson, err error) {
	c = &ContactPerson{
		contactPersonDo: q.ContactPerson.WithContext(context.Background()),
		accountInfo:     acc,
		wechat:          w,
	}
	for _, info := range w.WatchInfo {
		db, err := sqlite.NewSQLite(w.Key, path.Join(info.Path, "Contact"), dbName)
		if err != nil {
			return nil, err
		}
		query := contact.Use(db.DB())
		c.sqlInfo = append(c.sqlInfo, &sqliteInfo{
			query:       query,
			wcContactDo: query.WCContact.WithContext(context.Background()),
			id:          info.Id,
			hash:        info.Hash,
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
	contacts, err := info.wcContactDo.Find()
	if err != nil {
		return fmt.Errorf("accountId(%d) wcContactDo.Find(): %w", accId, err)
	}
	todo := make([]*model.ContactPerson, 0, len(contacts))
	for _, v := range contacts {
		cp := &model.ContactPerson{
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

	err = c.contactPersonDo.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(todo, 2000)
	if err != nil {
		return fmt.Errorf("contactPersonDo.CreateInBatches(): %w", err)
	}

	return nil
}
