package message

import (
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	"github.com/HarryWang29/echo_mind/internal/infra/db/sqlite/session"
)

type Message struct {
	messageDo              repo.IMessageDo
	sessionAbstractDo      repo.ISessionAbstractDo
	sessionAbstractBrandDo repo.ISessionAbstractBrandDo
	sessionSqlite          []*sessionInfo
}

type sessionInfo struct {
	query *session.Query

	id   string
	hash string
}
