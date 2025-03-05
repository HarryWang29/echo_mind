package main

import (
	"github.com/HarryWang29/echo_mind/config"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite"
	"gorm.io/gen"
	"path"
)

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error) // returns struct and error
}

func genDbSource(cfg *config.Config) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/infra/db/repo",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, err := repo.NewRepo(cfg.DataSource)
	if err != nil {
		panic(err)
	}
	g.UseDB(db) // reuse your gorm db

	ai := g.GenerateModel("account_info")
	gc := g.GenerateModel("group_contact")
	gcp := g.GenerateModel("group_contact_person")
	gur := g.GenerateModel("group_user_relation")
	cp := g.GenerateModel("contact_person")
	m := g.GenerateModel("message")
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(gc, gcp, cp, gur, m, ai)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, gc, gcp, cp, gur, m, ai)

	// Generate the code
	g.Execute()
}

func genSqliteSource(cfg *config.Config) {
	wc := config.GetWechatConfig(cfg)
	outPath := "internal/infra/db/sqlite"
	genSqlite(path.Join(outPath, "contact"), wc.Key, path.Join(wc.WatchInfo[0].Path, "Contact"), "wccontact_new2.db", "WCContact")
	genSqlite(path.Join(outPath, "group"), wc.Key, path.Join(wc.WatchInfo[0].Path, "Group"), "group_new.db",
		"GroupContact", "GroupMember", "GroupUserRelation",
	)
}

func genSqlite(outPath, key, dbPath, dbName string, tables ...string) {
	if len(tables) == 0 {
		return
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: outPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, err := sqlite.NewSQLite(key, dbPath, dbName)
	if err != nil {
		panic(err)
	}
	g.UseDB(db.DB()) // reuse your gorm db
	todo := make([]interface{}, len(tables))

	for i, table := range tables {
		t := g.GenerateModel(table)
		todo[i] = t
	}
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(todo...)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, todo...)

	// Generate the code
	g.Execute()
}

func main() {
	cfg := config.NewConfig()
	genDbSource(cfg)
	genSqliteSource(cfg)
}
