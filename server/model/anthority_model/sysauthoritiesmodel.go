package anthority_model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysAuthoritiesModel = (*customSysAuthoritiesModel)(nil)

type (
	// SysAuthoritiesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysAuthoritiesModel.
	SysAuthoritiesModel interface {
		sysAuthoritiesModel
		GetAllAuthorities(ctx context.Context) ([]*SysAuthorities, error)
		GetAuthorityUser(ctx context.Context, userId int64) ([]*SysAuthorities, error)
	}

	customSysAuthoritiesModel struct {
		*defaultSysAuthoritiesModel
	}
)

// NewSysAuthoritiesModel returns a model for the database table.
func NewSysAuthoritiesModel(conn sqlx.SqlConn) SysAuthoritiesModel {
	return &customSysAuthoritiesModel{
		defaultSysAuthoritiesModel: newSysAuthoritiesModel(conn),
	}
}

func (m *defaultSysAuthoritiesModel) GetAllAuthorities(ctx context.Context) ([]*SysAuthorities, error) {
	query := fmt.Sprintf("select * from %s", m.table)
	authorities := make([]*SysAuthorities, 0)
	err := m.conn.QueryRowsCtx(ctx, &authorities, query)
	if err != nil {
		return nil, err
	}
	return authorities, nil
}

// 根据用户ID 查找对应的权限列表
func (m *defaultSysAuthoritiesModel) GetAuthorityUser(ctx context.Context, userId int64) ([]*SysAuthorities, error) {
	query := fmt.Sprintf("SELECT * from  %s WHERE authority_id in (SELECT sys_authority_authority_id from `sys_user_authority`  WHERE sys_user_id = ?)", m.table)
	authorities := make([]*SysAuthorities, 0)
	err := m.conn.QueryRowsCtx(ctx, &authorities, query, userId)
	if err != nil {
		return nil, err
	}
	return authorities, nil
}
