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
		Inserts(datas []*SysUserAuthority) error
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
			return err
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

//创建用户批量插入

func (m *defaultSysUserAuthorityModel) Inserts(datas []*SysUserAuthority) error {

	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserAuthorityRowsExpectAutoSet)
	blk, err := sqlx.NewBulkInserter(m.conn, query)
	if err != nil {
		return err
	}
	defer blk.Flush()
	for _, data := range datas {
		err := blk.Insert(data.SysUserId, data.SysAuthorityAuthorityId)
		if err != nil {
			return err
		}
	}
	return nil
}
