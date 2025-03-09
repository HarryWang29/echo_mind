package group

import (
    "fmt"
    "github.com/HarryWang29/echo_mind/internal/infra/db/model"
    "github.com/HarryWang29/echo_mind/pkg/util"
    "gorm.io/gorm/clause"
)

func (c *Group) SyncContactPerson(account *model.AccountInfo) error {
    defer util.FuncCost(util.FuncName())()
    err := c.syncContactPerson(account.ID)
    if err != nil {
        return fmt.Errorf("sync account(%s) contact person: %w", c.sqlInfo.hash, err)
    }
    return nil
}

func (c *Group) syncContactPerson(accId int64) error {
    members, err := c.sqlInfo.contactPersonDo.Find()
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
