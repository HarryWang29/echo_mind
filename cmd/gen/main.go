package main

import (
	"github.com/HarryWang29/echo_mind/config"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm/schema"
	"path"
	"regexp"
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
	sa := g.GenerateModel("session_abstract")
	sab := g.GenerateModel("session_abstract_brand")
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(gc, gcp, cp, gur, m, ai, sa, sab)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, gc, gcp, cp, gur, m, ai)

	// Generate the code
	g.Execute()
}

func genSqliteSource(cfg *config.Config) {
	wconfig := config.GetWechatConfig(cfg)
	wc := wconfig.WatchInfo[0]
	outPath := "internal/infra/db/sqlite"
	genSqlite(path.Join(outPath, "contact"), wc.Key, path.Join(wc.Path, "Contact"), "wccontact_new2.db", "WCContact")
	genSqlite(path.Join(outPath, "group"), wc.Key, path.Join(wc.Path, "Group"), "group_new.db",
		"GroupContact", "GroupMember", "GroupUserRelation",
	)
	genSqlite(path.Join(outPath, "session"), wc.Key, path.Join(wc.Path, "Session"), "session_new.db",
		"SessionAbstract", "SessionAbstractBrand",
	)
	genSqlite(path.Join(outPath, "message"), wc.Key, path.Join(wc.Path, "Message"), "msg_0.db", "Chat")
}

type Message struct{}

func (m Message) TableName(namer schema.Namer) string {
	return "Chat"
}

func genSqlite(outPath, key, dbPath, dbName string, tables ...string) {
	if len(tables) == 0 {
		return
	}
	cfg := gen.Config{
		OutPath: outPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}
	cfg.WithFileNameStrategy(func(tableName string) (fileName string) {
		pattern := `^Chat_[A-Za-z0-9]+$`
		re := regexp.MustCompile(pattern)
		if re.MatchString(tableName) {
			return "message"
		}
		return tableName
	})
	g := gen.NewGenerator(cfg)

	db, err := sqlite.NewSQLite(key, dbPath, dbName)
	if err != nil {
		panic(err)
	}
	g.UseDB(db.DB()) // reuse your gorm db
	todo := make([]interface{}, 0, len(tables))

	if tables[0] == "Chat" {
		tableList, err := db.DB().Migrator().GetTables()
		if err != nil {
			panic(err)
		}
		pattern := `^Chat_[A-Za-z0-9]+$`
		re := regexp.MustCompile(pattern)
		for _, table := range tableList {
			if re.MatchString(table) {
				t := g.GenerateModelAs(table, "Message", gen.WithMethod(Message{}.TableName),
					gen.FieldGenType("mesSvrID", "Int64"))
				todo = append(todo, t)
				break
			}
		}

	} else {
		for _, table := range tables {
			t := g.GenerateModel(table)
			todo = append(todo, t)
		}
	}
	if len(todo) == 0 {
		return
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
