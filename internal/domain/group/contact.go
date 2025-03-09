package group

import (
    "fmt"
    "github.com/HarryWang29/echo_mind/internal/infra/db/model"
    "github.com/HarryWang29/echo_mind/pkg/util"
    "gorm.io/gorm/clause"
)

func (c *Group) SyncContact(account *model.AccountInfo) error {
    defer util.FuncCost(util.FuncName())()
    err := c.syncContact(account.ID)
    if err != nil {
        return fmt.Errorf("sync account(%s) contact person: %w", c.sqlInfo.hash, err)
    }
    return nil
}

func (c *Group) syncContact(accId int64) error {
    contacts, err := c.sqlInfo.contactDo.Find()
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
