package group

import (
	"fmt"
	"github.com/HarryWang29/echo_mind/internal/infra/db/model"
	"github.com/HarryWang29/echo_mind/pkg/util"
	"gorm.io/gorm/clause"
)

func (c *Group) SyncContactPerson() error {
	for _, info := range c.sqlInfo {
		account, err := c.accountInfo.FindByHash(info.hash)
		if err != nil {
			return fmt.Errorf("find account(%s) info by hash: %w", info.hash, err)
		}
		err = c.syncContactPerson(account.ID, info)
		if err != nil {
			return fmt.Errorf("sync account(%s) contact person: %w", info.hash, err)
		}
	}
	return nil
}

func (c *Group) syncContactPerson(accId int64, info *sqliteInfo) error {
	members, err := info.contactPersonDo.Find()
	if err != nil {
		return fmt.Errorf("accountId(%d) contactPersonDo.Find(): %w", accId, err)
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

	err = c.contactPersonDo.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(todo, 2000)
	if err != nil {
		return fmt.Errorf("contactPersonDo.CreateInBatches(): %w", err)
	}
	return nil
}
