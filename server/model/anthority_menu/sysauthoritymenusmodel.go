package anthority_menu

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysAuthorityMenusModel = (*customSysAuthorityMenusModel)(nil)

type (
	// SysAuthorityMenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysAuthorityMenusModel.
	SysAuthorityMenusModel interface {
		sysAuthorityMenusModel
	}

	customSysAuthorityMenusModel struct {
		*defaultSysAuthorityMenusModel
	}
)

// NewSysAuthorityMenusModel returns a model for the database table.
func NewSysAuthorityMenusModel(conn sqlx.SqlConn) SysAuthorityMenusModel {
	return &customSysAuthorityMenusModel{
		defaultSysAuthorityMenusModel: newSysAuthorityMenusModel(conn),
	}
}
