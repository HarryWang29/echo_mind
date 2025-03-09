package message

import (
    "github.com/HarryWang29/echo_mind/config"
    "github.com/HarryWang29/echo_mind/internal/infra/db/repo"
    "github.com/HarryWang29/echo_mind/pkg/util"
    "testing"
)

func TestTableName(t *testing.T) {
    cfg := config.NewConfig()
    wconfig := config.GetWechatConfig(cfg)
    dataSourceConfig := config.GetDataSourceConfig(cfg)
    db, err := repo.NewRepo(dataSourceConfig)
    if err != nil {
        t.Fatal(err)
    }
    v := repo.GetOptions()
    query := repo.Use(db, v...)
    m, err := New(&wconfig.WatchInfo[0], query)
    if err != nil {
        t.Fatal(err)
    }
    tableName := "02bdd9a2e3ecf8c42207ebbac761f7ce"
    q := m.msgDbs[tableName]
    find, err := q.do.Find()
    if err != nil {
        t.Fatal(err)
    }
    for _, message := range find {
        t.Log(message.MsgContent)
    }
}

func TestFind(t *testing.T) {
    cfg := config.NewConfig()
    wconfig := config.GetWechatConfig(cfg)
    dataSourceConfig := config.GetDataSourceConfig(cfg)
    db, err := repo.NewRepo(dataSourceConfig)
    if err != nil {
        t.Fatal(err)
    }
    v := repo.GetOptions()
    query := repo.Use(db, v...)
    m, err := New(&wconfig.WatchInfo[0], query)
    if err != nil {
        t.Fatal(err)
    }
    count, err := m.messageDo.
        Where(
            query.Message.AccountID.Eq(1),
            query.Message.Hash.Eq(util.HashHex(util.MD5, "")),
        ).Count()
    if err != nil {
        t.Fatal(err)
    }
    t.Log(count)
}
