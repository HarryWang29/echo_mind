package group

import (
	"fmt"
	"github.com/HarryWang29/echo_mind/internal/infra/db/model"
	"github.com/HarryWang29/echo_mind/pkg/util"
	"gorm.io/gorm/clause"
)

func (c *Group) SyncContact() error {
	for _, info := range c.sqlInfo {
		account, err := c.accountInfo.FindByHash(info.hash)
		if err != nil {
			return fmt.Errorf("find account(%s) info by hash: %w", info.hash, err)
		}
		err = c.syncContact(account.ID, info)
		if err != nil {
			return fmt.Errorf("sync account(%s) contact person: %w", info.hash, err)
		}
	}
	return nil
}

func (c *Group) syncContact(accId int64, info *sqliteInfo) error {
	contacts, err := info.contactDo.Find()
	if err != nil {
		return fmt.Errorf("accountId(%d) contactDo.Find(): %w", accId, err)
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

	err = c.contactDo.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(todo, 2000)
	if err != nil {
		return fmt.Errorf("contactDo.CreateInBatches(): %w", err)
	}
	return nil
}
