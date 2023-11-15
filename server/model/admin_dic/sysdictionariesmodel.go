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
		GetDicList(ctx context.Context, req *types.GetDicListReq) ([]*SysDictionaries, int, error)
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

func (m *defaultSysDictionariesModel) GetDicList(ctx context.Context, req *types.GetDicListReq) ([]*SysDictionaries, int, error) {
	var DicArr []*SysDictionaries
	total := 0
	if req.Page < 1 {
		req.Page = 1
	}
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	if req.Name != "" {
		query := fmt.Sprintf("select * from %s where name like ?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, "%"+req.Name+"%")
		total = len(DicArr)
		if err != nil {
			return nil, 0, err
		}
		return DicArr, total, nil

	}
	if req.Type != "" {
		query := fmt.Sprintf("select * from %s where type like ?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, "%"+req.Type+"%")
		total = len(DicArr)
		if err != nil {
			return nil, 0, err
		}
		return DicArr, total, nil

	}

	if req.Status != 0 {
		query := fmt.Sprintf("select * from %s where status =?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, req.Status)
		total = len(DicArr)
		if err != nil {
			return nil, 0, err
		}
		return DicArr, total, nil

	}

	if req.Desc != "" {
		query := fmt.Sprintf("select * from %s where desc like ?", m.table)
		err := m.conn.QueryRowsCtx(ctx, &DicArr, query, "%"+req.Desc+"%")
		total = len(DicArr)
		if err != nil {
			return nil, 0, err
		}
		return DicArr, total, nil

	}

	queryCount := fmt.Sprintf("select COUNT(1) FROM %s ", m.table)
	err := m.conn.QueryRow(&total, queryCount)
	if err != nil {
		return nil, 0, err
	}
	query := fmt.Sprintf("select * from %s limit ?,?", m.table)
	err = m.conn.QueryRowsCtx(ctx, &DicArr, query, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return DicArr, total, nil
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
