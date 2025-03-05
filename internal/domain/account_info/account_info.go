package account_info

import (
	"context"
	"fmt"
	"github.com/HarryWang29/echo_mind/internal/infra/db/model"
	"github.com/HarryWang29/echo_mind/internal/infra/db/repo"
)

type AccountInfo struct {
	query         *repo.Query
	accountInfoDo repo.IAccountInfoDo
}

func NewAccountInfo(q *repo.Query) (a *AccountInfo, err error) {
	a = &AccountInfo{
		query:         q,
		accountInfoDo: q.AccountInfo.WithContext(context.Background()),
	}
	return a, nil
}

func (a *AccountInfo) FindByHash(hash string) (*model.AccountInfo, error) {
	info, err := a.accountInfoDo.Where(a.query.AccountInfo.Hash.Eq(hash)).FirstOrCreate()
	if err != nil {
		return nil, fmt.Errorf("find account info by hash: %w", err)
	}
	return info, nil
}
