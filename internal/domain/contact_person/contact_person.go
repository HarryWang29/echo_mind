package contact_person

import (
    "context"
    "fmt"
    "github.com/HarryWang29/echo_mind/config"
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
    wechat          *config.WechatWatchInfo
    sqlInfo         *sqliteInfo
}

type sqliteInfo struct {
    query       *contact.Query
    wcContactDo contact.IWCContactDo
    id          string
    hash        string
}

const dbName = "wccontact_new2.db"

func NewContactPerson(w *config.WechatWatchInfo, q *repo.Query) (c *ContactPerson, err error) {
    c = &ContactPerson{
        contactPersonDo: q.ContactPerson.WithContext(context.Background()),
        wechat:          w,
    }
    db, err := sqlite.NewSQLite(w.Key, path.Join(w.Path, "Contact"), dbName)
    if err != nil {
        return nil, err
    }
    query := contact.Use(db.DB())
    c.sqlInfo = &sqliteInfo{
        query:       query,
        wcContactDo: query.WCContact.WithContext(context.Background()),
        id:          w.Id,
        hash:        w.Hash,
    }
    return c, nil
}

func (c *ContactPerson) Sync(account *model.AccountInfo) error {
    defer util.FuncCost(util.FuncName())()
    err := c.SyncContactPerson(account.ID)
    if err != nil {
        return fmt.Errorf("sync account(%s) contact person: %w", c.sqlInfo.hash, err)
    }
    return nil
}

func (c *ContactPerson) SyncContactPerson(accId int64) error {
    contacts, err := c.sqlInfo.wcContactDo.Find()
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
