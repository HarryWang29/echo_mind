package message

import (
    "context"
    "fmt"
    "github.com/HarryWang29/echo_mind/config"
    "github.com/HarryWang29/echo_mind/internal/infra/db/repo"
    "github.com/HarryWang29/echo_mind/internal/infra/db/sqlite"
    "github.com/HarryWang29/echo_mind/internal/infra/db/sqlite/message"
    "github.com/HarryWang29/echo_mind/internal/infra/db/sqlite/session"
    "os"
    "path"
    "path/filepath"
    "strings"
)

type Message struct {
    query                  *repo.Query
    messageDo              repo.IMessageDo
    sessionAbstractDo      repo.ISessionAbstractDo
    sessionAbstractBrandDo repo.ISessionAbstractBrandDo
    sessionAbstract        session.ISessionAbstractDo
    sessionAbstractBrand   session.ISessionAbstractBrandDo
    wechat                 *config.WechatWatchInfo
    msgDbs                 map[string]*msgDb
}

type msgDb struct {
    do     message.IMessageDo
    dbName string
    query  *message.Query
}

const sessionDbName = "session_new.db"

func New(w *config.WechatWatchInfo, q *repo.Query) (m *Message, err error) {
    m = &Message{
        query:                  q,
        messageDo:              q.Message.Table("message").WithContext(context.Background()),
        sessionAbstractDo:      q.SessionAbstract.WithContext(context.Background()),
        sessionAbstractBrandDo: q.SessionAbstractBrand.WithContext(context.Background()),
        wechat:                 w,
        msgDbs:                 make(map[string]*msgDb),
    }
    msgDbDir := path.Join(m.wechat.Path, "Message")
    msgDbName, err := m.getMsgDbPath(msgDbDir)
    if err != nil {
        return nil, fmt.Errorf("getMsgDbPath: %w", err)
    }
    for _, dbName := range msgDbName {
        db, err := sqlite.NewSQLite(w.Key, msgDbDir, dbName)
        if err != nil {
            return nil, fmt.Errorf("sqlite.NewSQLite: %w", err)
        }
        tables, err := db.DB().Migrator().GetTables()
        if err != nil {
            return nil, fmt.Errorf("db.Migrator().GetTables(): %w", err)
        }
        for _, table := range tables {
            if strings.HasPrefix(table, "Chat_") && !strings.HasSuffix(table, "_dels") {
                hash := strings.TrimPrefix(table, "Chat_")
                query := message.Use(db.DB())
                dbInfo := &msgDb{
                    do:     query.Message.Table(table).WithContext(context.Background()),
                    dbName: dbName,
                    query:  query,
                }
                m.msgDbs[hash] = dbInfo
            }
        }
    }
    db, err := sqlite.NewSQLite(w.Key, path.Join(m.wechat.Path, "Session"), sessionDbName)
    if err != nil {
        return nil, fmt.Errorf("sqlite.NewSQLite: %w", err)
    }
    query := session.Use(db.DB())
    m.sessionAbstract = query.SessionAbstract.WithContext(context.Background())
    m.sessionAbstractBrand = query.SessionAbstractBrand.WithContext(context.Background())

    return m, nil
}

func (m *Message) getMsgDbPath(msgDbDir string) (msgDbName []string, err error) {
    err = filepath.Walk(msgDbDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasPrefix(info.Name(), "msg_") && filepath.Ext(path) == ".db" {
            msgDbName = append(msgDbName, info.Name())
        }
        return nil
    })
    if err != nil {
        return nil, fmt.Errorf("walk msgDbDir(%s): %w", msgDbDir, err)
    }
    return msgDbName, nil
}
