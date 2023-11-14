package admin_dic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-server/server/internal/types"
	"zero-server/server/model/admin_dic_detail"
)

var _ SysDictionariesModel = (*customSysDictionariesModel)(nil)

type (
	// SysDictionariesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictionariesModel.
	SysDictionariesModel interface {
		sysDictionariesModel
		GetDicList(ctx context.Context, req *types.GetDicListReq) ([]*SysDictionaries, error)
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

func (m *defaultSysDictionariesModel) GetDicList(ctx context.Context, req *types.GetDicListReq) ([]*SysDictionaries, error) {
	var DicArr []*SysDictionaries
	if req.Name != "" {
		query := fmt.Sprintf("select * from %s where name =?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, req.Name)
		if err != nil {
			return nil, err
		}
		return DicArr, nil

	}
	if req.Type != "" {
		query := fmt.Sprintf("select * from %s where type =?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, req.Type)
		if err != nil {
			return nil, err
		}
		return DicArr, nil

	}

	if req.Status != 0 {
		query := fmt.Sprintf("select * from %s where status =?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, req.Status)
		if err != nil {
			return nil, err
		}
		return DicArr, nil

	}

	if req.Desc != "" {
		query := fmt.Sprintf("select * from %s where desc =?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, req.Desc)
		if err != nil {
			return nil, err
		}
		return DicArr, nil

	}
	query := fmt.Sprintf("select * from %s", m.table)
	err := m.conn.QueryRowsCtx(ctx, &DicArr, query)
	if err != nil {
		return nil, err
	}
	return DicArr, nil
}

func (m *defaultSysDictionariesModel) GetDicDetails(ctx context.Context, id int64) ([]*admin_dic_detail.SysDictionaryDetails, error) {
	var DicDetailsArr []*admin_dic_detail.SysDictionaryDetails

	query := fmt.Sprintf("select * from sys_dictionary_details where sys_dictionary_id =?")
	err := m.conn.QueryRowsCtx(ctx, &DicDetailsArr, query, id)
	if err != nil {
		return nil, err
	}
	return DicDetailsArr, nil
}
