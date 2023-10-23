package user_auth

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserAuthorityModel = (*customSysUserAuthorityModel)(nil)

type (
	// SysUserAuthorityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserAuthorityModel.
	SysUserAuthorityModel interface {
		sysUserAuthorityModel
		SetUserAuth(ctx context.Context, authIds []int64, userId int64) error
	}

	customSysUserAuthorityModel struct {
		*defaultSysUserAuthorityModel
	}
)

// NewSysUserAuthorityModel returns a model for the database table.
func NewSysUserAuthorityModel(conn sqlx.SqlConn) SysUserAuthorityModel {
	return &customSysUserAuthorityModel{
		defaultSysUserAuthorityModel: newSysUserAuthorityModel(conn),
	}
}

func (m *defaultSysUserAuthorityModel) SetUserAuth(ctx context.Context, authIds []int64, userId int64) error {
	err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		query := fmt.Sprintf("delete from sys_user_authority WHERE sys_user_id = ? ")
		_, err := m.conn.ExecCtx(ctx, query, userId)
		if err != nil {
			return err
		}
		//批量插入
		query = fmt.Sprintf("INSERT into sys_user_authority (sys_authority_authority_id,sys_user_id) VALUES(?,?)")
		blk, err := sqlx.NewBulkInserter(m.conn, query)
		if err != nil {
			panic(err)
		}
		defer blk.Flush()

		for _, authId := range authIds {
			err := blk.Insert(authId, userId)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
