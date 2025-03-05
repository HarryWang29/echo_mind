package sync_db

import (
	"fmt"
	"github.com/HarryWang29/echo_mind/config"
	"github.com/HarryWang29/echo_mind/internal/domain/contact_person"
	"github.com/HarryWang29/echo_mind/internal/domain/group"
	"github.com/HarryWang29/echo_mind/pkg/util"
)

type App struct {
	wechat        *config.WechatConfig
	contactPerson *contact_person.ContactPerson
	group         *group.Group
}

func NewApp(w *config.WechatConfig,
	cp *contact_person.ContactPerson,
	group *group.Group,
) (a *App, err error) {
	a = &App{
		wechat:        w,
		contactPerson: cp,
		group:         group,
	}
	return a, nil
}

func (app *App) Sync() (err error) {
	defer util.FuncCost(util.FuncName())()
	err = app.contactPerson.Sync()
	if err != nil {
		return fmt.Errorf("contactPerson.Sync(): %w", err)
	}

	err = app.group.SyncContact()
	if err != nil {
		return fmt.Errorf("groupContact.Sync(): %w", err)
	}

	err = app.group.SyncContactPerson()
	if err != nil {
		return fmt.Errorf("groupContactPerson.Sync(): %w", err)
	}

	err = app.group.SyncUserRelation()
	if err != nil {
		return fmt.Errorf("groupContactPerson.Sync(): %w", err)
	}
	return nil
}
