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
		GetOperationList(ctx context.Context) ([]*SysOperationRecords, error)
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

func (m *defaultSysOperationRecordsModel) GetOperationList(ctx context.Context) ([]*SysOperationRecords, error) {
	query := fmt.Sprintf("select * from %s", m.table)
	operations := make([]*SysOperationRecords, 0)
	err := m.conn.QueryRowsCtx(ctx, &operations, query)
	return operations, err
}
