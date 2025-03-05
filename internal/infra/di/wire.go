//go:build wireinject
// +build wireinject

package di

import (
	"github.com/HarryWang29/echo_mind/config"
	sync_db "github.com/HarryWang29/echo_mind/internal/application/sync"
	"github.com/HarryWang29/echo_mind/internal/domain/account_info"
	"github.com/HarryWang29/echo_mind/internal/domain/contact_person"
	"github.com/HarryWang29/echo_mind/internal/domain/group"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	"github.com/google/wire"
)

func InjectAll() (*sync_db.App, error) {
	wire.Build(
		config.NewConfig,
		config.GetWechatConfig,
		config.GetDataSourceConfig,
		repo.GetOptions,
		repo.NewRepo,
		repo.Use,
		account_info.NewAccountInfo,
		contact_person.NewContactPerson,
		group.New,
		sync_db.NewApp,
	)
	return &sync_db.App{}, nil
}
