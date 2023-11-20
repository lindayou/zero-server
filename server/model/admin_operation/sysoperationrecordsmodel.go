package admin_operation

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysOperationRecordsModel = (*customSysOperationRecordsModel)(nil)

type (
	// SysOperationRecordsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysOperationRecordsModel.
	SysOperationRecordsModel interface {
		sysOperationRecordsModel
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
