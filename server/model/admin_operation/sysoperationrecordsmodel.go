package admin_operation

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysOperationRecordsModel = (*customSysOperationRecordsModel)(nil)

type (
	// SysOperationRecordsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysOperationRecordsModel.
	SysOperationRecordsModel interface {
		sysOperationRecordsModel
		GetOperationList(ctx context.Context, PageSize int, Offset int) ([]*SysOperationRecords, int, error)
	}

	customSysOperationRecordsModel struct {
		*defaultSysOperationRecordsModel
	}
)

// NewSysOperationRecordsModel returns a model for the database table.
func NewSysOperationRecordsModel(conn sqlx.SqlConn) SysOperationRecordsModel {
	return &customSysOperationRecordsModel{
		defaultSysOperationRecordsModel: newSysOperationRecordsModel(conn),
	}
}

func (m *defaultSysOperationRecordsModel) GetOperationList(ctx context.Context, PageSize int, Offset int) ([]*SysOperationRecords, int, error) {
	query := fmt.Sprintf("select * from %s order by created_at desc limit ?,? ", m.table)
	operations := make([]*SysOperationRecords, 0)
	err := m.conn.QueryRowsCtx(ctx, &operations, query, Offset, PageSize)
	if err != nil {
		return nil, 0, err
	}

	var total = 0
	query2 := fmt.Sprintf("select count(1) from %s", m.table)
	err = m.conn.QueryRowCtx(ctx, &total, query2)
	if err != nil {
		return nil, 0, err
	}
	return operations, total, err
}
