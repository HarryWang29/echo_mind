package group

import (
	"fmt"
	"github.com/HarryWang29/echo_mind/internal/infra/db/model"
	"gorm.io/gorm/clause"
)

func (c *Group) SyncUserRelation(account *model.AccountInfo) error {
	err := c.syncUserRelation(account.ID)
	if err != nil {
		return fmt.Errorf("sync account(%s) contact person: %w", c.sqlInfo.hash, err)
	}
	return nil
}

func (c *Group) syncUserRelation(accId int64) error {
	contacts, err := c.sqlInfo.userRelation.Find()
	if err != nil {
		return fmt.Errorf("accountId(%d) userRelation.Find(): %w", accId, err)
	}
	todo := make([]*model.GroupUserRelation, 0, len(contacts))
	for _, v := range contacts {
		cp := &model.GroupUserRelation{
			AccountID: accId,
			UserName:  v.UserName,
			GroupName: v.GroupNameList,
		}
		todo = append(todo, cp)
	}

	err = c.userRelation.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(todo, 2000)
	if err != nil {
		return fmt.Errorf("userRelation.CreateInBatches(): %w", err)
	}
	return nil
}
