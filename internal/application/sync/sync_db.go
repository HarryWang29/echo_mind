package sync_db

import (
    "fmt"
    "github.com/HarryWang29/echo_mind/config"
    "github.com/HarryWang29/echo_mind/internal/domain/account_info"
    "github.com/HarryWang29/echo_mind/internal/domain/contact_person"
    "github.com/HarryWang29/echo_mind/internal/domain/group"
    "github.com/HarryWang29/echo_mind/internal/domain/message"
    "github.com/HarryWang29/echo_mind/internal/infra/db/repo"
    "github.com/HarryWang29/echo_mind/pkg/util"
)

type App struct {
    wechat      *config.WechatConfig
    accountInfo *account_info.AccountInfo
    tasks       []*taskInfo
}

type taskInfo struct {
    contactPerson *contact_person.ContactPerson
    group         *group.Group
    watchInfo     *config.WechatWatchInfo
    message       *message.Message
}

func NewApp(w *config.WechatConfig,
    query *repo.Query,
    acc *account_info.AccountInfo,
) (a *App, err error) {
    a = &App{
        wechat:      w,
        accountInfo: acc,
    }
    for _, info := range w.WatchInfo {
        contactPerson, err := contact_person.NewContactPerson(&info, query)
        if err != nil {
            return nil, fmt.Errorf("new contact person: %w", err)
        }
        group, err := group.New(&info, query)
        if err != nil {
            return nil, fmt.Errorf("new group: %w", err)
        }
        m, err := message.New(&info, query)
        if err != nil {
            return nil, fmt.Errorf("new message: %w", err)
        }
        a.tasks = append(a.tasks, &taskInfo{
            contactPerson: contactPerson,
            group:         group,
            watchInfo:     &info,
            message:       m,
        })
    }
    return a, nil
}

func (app *App) Sync() (err error) {
    defer util.FuncCost(util.FuncName())()
    for _, task := range app.tasks {
        account, err := app.accountInfo.FindByHash(task.watchInfo.Hash)
        if err != nil {
            return fmt.Errorf("find account(%s) info by hash: %w", task.watchInfo.Hash, err)
        }
        err = task.contactPerson.Sync(account)
        if err != nil {
            return fmt.Errorf("contactPerson.Sync(): %w", err)
        }

        err = task.group.SyncContact(account)
        if err != nil {
            return fmt.Errorf("groupContact.Sync(): %w", err)
        }

        err = task.group.SyncContactPerson(account)
        if err != nil {
            return fmt.Errorf("groupContactPerson.Sync(): %w", err)
        }

        err = task.group.SyncUserRelation(account)
        if err != nil {
            return fmt.Errorf("groupContactPerson.Sync(): %w", err)
        }

        err = task.message.Sync(account)
        if err != nil {
            return fmt.Errorf("message.Sync(): %w", err)
        }
    }
    return nil
}
