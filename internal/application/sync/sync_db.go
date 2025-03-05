package sync_db

import (
	"fmt"
	"github.com/HarryWang29/echo_mind/config"
	"github.com/HarryWang29/echo_mind/internal/domain/contact_person"
	"github.com/HarryWang29/echo_mind/internal/domain/group"
)

type App struct {
	wechat             *config.WechatConfig
	contactPerson      *contact_person.ContactPerson
	groupContact       *group.Contact
	groupContactPerson *group.ContactPerson
}

func NewApp(w *config.WechatConfig,
	cp *contact_person.ContactPerson,
	gc *group.Contact,
	gcp *group.ContactPerson,
) (a *App, err error) {
	a = &App{
		wechat:             w,
		contactPerson:      cp,
		groupContact:       gc,
		groupContactPerson: gcp,
	}
	return a, nil
}

func (app *App) Sync() (err error) {
	err = app.contactPerson.Sync()
	if err != nil {
		return fmt.Errorf("contactPerson.Sync(): %w", err)
	}

	err = app.groupContact.Sync()
	if err != nil {
		return fmt.Errorf("groupContact.Sync(): %w", err)
	}

	err = app.groupContactPerson.Sync()
	if err != nil {
		return fmt.Errorf("groupContactPerson.Sync(): %w", err)
	}
	return nil
}
