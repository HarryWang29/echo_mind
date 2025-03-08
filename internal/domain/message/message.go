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
	messageDo              repo.IMessageDo
	sessionAbstractDo      repo.ISessionAbstractDo
	sessionAbstractBrandDo repo.ISessionAbstractBrandDo
	sessionSqlite          *sessionInfo
	wechat                 *config.WechatWatchInfo
	msgDbs                 map[string]message.IMessageDo
}

type sessionInfo struct {
	query *session.Query

	id   string
	hash string
}

func New(w *config.WechatWatchInfo, q *repo.Query) (m *Message, err error) {
	m = &Message{
		messageDo:              q.Message.WithContext(context.Background()),
		sessionAbstractDo:      q.SessionAbstract.WithContext(context.Background()),
		sessionAbstractBrandDo: q.SessionAbstractBrand.WithContext(context.Background()),
		wechat:                 w,
		msgDbs:                 make(map[string]message.IMessageDo),
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
				query.Message.UseTable(table)
				m.msgDbs[hash] = query.Message.WithContext(context.Background())
			}
		}
	}
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
