package admin_dic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysDictionariesModel = (*customSysDictionariesModel)(nil)

type (
	// SysDictionariesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictionariesModel.
	SysDictionariesModel interface {
		sysDictionariesModel
		GetDicList(ctx context.Context) ([]*SysDictionaries, error)
	}

	customSysDictionariesModel struct {
		*defaultSysDictionariesModel
	}
)

// NewSysDictionariesModel returns a model for the database table.
func NewSysDictionariesModel(conn sqlx.SqlConn) SysDictionariesModel {
	return &customSysDictionariesModel{
		defaultSysDictionariesModel: newSysDictionariesModel(conn),
	}
}

func (m *defaultSysDictionariesModel) GetDicList(ctx context.Context) ([]*SysDictionaries, error) {
	var DicArr []*SysDictionaries
	query := fmt.Sprintf("select * from %s", m.table)
	err := m.conn.QueryRowsCtx(ctx, &DicArr, query)
	if err != nil {
		return nil, err
	}
	return DicArr, nil
}
