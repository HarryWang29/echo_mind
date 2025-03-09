package message

import (
	"fmt"
	sourceModel "github.com/HarryWang29/echo_mind/internal/infra/db/model"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
	localModel "github.com/HarryWang29/echo_mind/internal/infra/db/sqlite/model"
	"github.com/HarryWang29/echo_mind/pkg/util"
	"gorm.io/gorm/clause"
)

func (m *Message) Sync(account *sourceModel.AccountInfo) error {
	defer util.FuncCost(util.FuncName())()
	sessions, err := m.getSessions()
	if err != nil {
		return fmt.Errorf("get sessions: %v", err)
	}
	for _, session := range sessions {
		last, err := m.checkSourceLastMsg(account, session)
		if err != nil {
			return fmt.Errorf("check source last msg: %w", err)
		}
		if last == 0 {
			continue
		}
		err = m.syncMsg(account, session.MNsUserName, last)
		if err != nil {
			return fmt.Errorf("sync msg: %w", err)
		}
		break
	}
	return nil
}

func (m *Message) syncMsg(account *sourceModel.AccountInfo, name string, last int64) error {
	hash := util.HashHex(util.MD5, name)
	db, ok := m.msgDbs[hash]
	if !ok {
		return fmt.Errorf("msg db(%s) not found", hash)
	}
	msgs, err := db.do.Where(db.query.Message.MsgCreateTime.Gt(int32(last))).Find()
	if err != nil {
		return fmt.Errorf("find msg create time: %w", err)
	}
	if len(msgs) == 0 {
		return nil
	}
	todo := make([]*sourceModel.Message, 0, len(msgs))
	for _, msg := range msgs {
		todo = append(todo, &sourceModel.Message{
			AccountID:   account.ID,
			Hash:        hash,
			LocalID:     int64(msg.MesLocalID),
			SvrID:       msg.MesSvrID,
			CreateTime:  int64(msg.MsgCreateTime),
			Content:     msg.MsgContent,
			Status:      int64(msg.MsgStatus),
			ImgStatus:   int64(msg.MsgImgStatus),
			MessageType: msg.MessageType,
			Des:         msg.MesDes == 1,
			Source:      msg.MsgSource,
			VoiceText:   msg.MsgVoiceText,
			Seq:         int64(msg.MsgSeq),
			DbName:      db.dbName,
		})
	}
	err = m.messageDo.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(todo, 2000)
	if err != nil {
		return fmt.Errorf("create msg: %w", err)
	}
	return nil
}

func (m *Message) getSessions() ([]*localModel.SessionAbstract, error) {
	find, err := m.sessionAbstract.Find()
	if err != nil {
		return nil, fmt.Errorf("get session abstract: %w", err)
	}
	return find, nil
}

func (m *Message) checkSourceLastMsg(account *sourceModel.AccountInfo, session *localModel.SessionAbstract) (last int64, err error) {
	count, err := m.messageDo.
		Where(
			m.query.Message.AccountID.Eq(account.ID),
			m.query.Message.Hash.Eq(util.HashHex(util.MD5, session.MNsUserName)),
		).Count()
	if err != nil {
		return 0, fmt.Errorf("count message: %w", err)
	}
	if count == 0 {
		return -1, nil
	}
	lastMessage, err := m.messageDo.Where(
		m.query.Message.AccountID.Eq(account.ID),
		m.query.Message.Hash.Eq(util.HashHex(util.MD5, session.MNsUserName)),
	).Order(repo.Message.CreateTime).Last()
	if err != nil {
		return 0, fmt.Errorf("get last msg: %w", err)
	}
	if lastMessage.CreateTime < int64(session.MULastTime) {
		return lastMessage.CreateTime, nil
	}
	return 0, nil
}
