package admin_dic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-server/server/model/admin_dic_detail"
)

var _ SysDictionariesModel = (*customSysDictionariesModel)(nil)

type (
	// SysDictionariesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictionariesModel.
	SysDictionariesModel interface {
		sysDictionariesModel
		GetDicList(ctx context.Context) ([]*SysDictionaries, error)
		GetDicDetails(ctx context.Context, id int64) ([]*admin_dic_detail.SysDictionaryDetails, error)
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

func (m *defaultSysDictionariesModel) GetDicDetails(ctx context.Context, id int64) ([]*admin_dic_detail.SysDictionaryDetails, error) {
	var DicDetailsArr = []*admin_dic_detail.SysDictionaryDetails{}
	query := fmt.Sprintf("select * from sys_dictionary_details where sys_dictionary_id =?")
	err := m.conn.QueryRowsCtx(ctx, &DicDetailsArr, query, id)
	if err != nil {
		return nil, err
	}
	return DicDetailsArr, nil
}
