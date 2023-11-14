package admin_dic_detail

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysDictionaryDetailsModel = (*customSysDictionaryDetailsModel)(nil)

type (
	// SysDictionaryDetailsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictionaryDetailsModel.
	SysDictionaryDetailsModel interface {
		sysDictionaryDetailsModel
	}

	customSysDictionaryDetailsModel struct {
		*defaultSysDictionaryDetailsModel
	}
)

// NewSysDictionaryDetailsModel returns a model for the database table.
func NewSysDictionaryDetailsModel(conn sqlx.SqlConn) SysDictionaryDetailsModel {
	return &customSysDictionaryDetailsModel{
		defaultSysDictionaryDetailsModel: newSysDictionaryDetailsModel(conn),
	}
}
