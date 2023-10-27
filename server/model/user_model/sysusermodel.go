package user_model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		//查询用户是否存在
		Find(ctx context.Context, data *SysUser) ([]*SysUser, error)

		FindByUsername(ctx context.Context, data *SysUser) (*SysUser, error)
		//用户列表
		UserList(ctx context.Context, page, pageSize int64) ([]*UserMessage, error, int)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
	UserMessage struct {
		Id          int64     `db:"id"`           // 用户ID
		Username    string    `db:"username"`     // 用户名
		Phone       string    `db:"phone"`        // 手机号
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdateAt    time.Time `db:"update_at"`    // 更新时间
		AuthorityId int       `db:"authority_id"` // 权限id
		Email       string    `db:"email"`        // 邮箱
		Enable      int64     `db:"enable"`       // 禁用标识
		Uuid        string    `db:"uuid"`         // 用户uuid
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

// 查询用户是否存在
func (m *defaultSysUserModel) Find(ctx context.Context, data *SysUser) ([]*SysUser, error) {
	var resp []*SysUser
	query := fmt.Sprintf("select * from %s where username = ? ", m.table)
	err := m.conn.QueryRows(&resp, query, data.Username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

// 根据用户名查询单个用户信息
func (m *defaultSysUserModel) FindByUsername(ctx context.Context, data *SysUser) (*SysUser, error) {
	var resp SysUser
	query := fmt.Sprintf("select * from %s where username = ? ", m.table)
	err := m.conn.QueryRow(&resp, query, data.Username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultSysUserModel) UserList(ctx context.Context, page, pageSize int64) ([]*UserMessage, error, int) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var resp = []*UserMessage{}
	query := fmt.Sprintf("select * from %s limit ?,?", m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, offset, pageSize)
	count := len(resp)
	if err != nil {
		return nil, err, 0
	}
	return resp, nil, count

}
