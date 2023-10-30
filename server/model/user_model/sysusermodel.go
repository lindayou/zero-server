package user_model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"zero-server/server/utils"
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
		//删除用户
		DeleteUser(ctx context.Context, id int64) error
		//更新用户
		UpdateUser(ctx context.Context, data *SysUser) error
		//修改用户密码
		ChangePass(ctx context.Context, priPass string, newPass string, userId int64) error
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

func (m *defaultSysUserModel) DeleteUser(ctx context.Context, id int64) error {
	err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		query := fmt.Sprintf("delete from %s where id =?", m.table)
		_, err := m.conn.ExecCtx(ctx, query, id)
		if err != nil {
			return err
		}
		query2 := fmt.Sprintf("delete from sys_user_authority where sys_user_id = ?")
		_, err = m.conn.ExecCtx(ctx, query2, id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}
func (m *defaultSysUserModel) UpdateUser(ctx context.Context, data *SysUser) error {
	err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
		_, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Phone, data.AuthorityId, data.Email, data.Enable, data.Uuid, data.Id)

		query2 := fmt.Sprintf("delete from sys_user_authority where sys_user_id = ?")
		_, err = m.conn.ExecCtx(ctx, query2, data.Id)
		if err != nil {
			return err
		}
		query = fmt.Sprintf("INSERT into sys_user_authority (sys_authority_authority_id,sys_user_id) VALUES(?,?)")
		blk, err := sqlx.NewBulkInserter(m.conn, query)
		if err != nil {
			return err
		}
		defer blk.Flush()

		for _, authId := range data.AuthorityIds {
			err := blk.Insert(authId, data.Id)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (m *defaultSysUserModel) ChangePass(ctx context.Context, priPass string, newPass string, userId int64) error {

	//验证原密码
	query := fmt.Sprintf("select password from  %s where id = ?", m.table)
	realPass := ""
	err := m.conn.QueryRowCtx(ctx, &realPass, query, userId)
	if err != nil {
		return err
	}
	if !utils.BcryptCheck(priPass, realPass) {
		err = errors.New("原密码错误")
		return err
	}

	//修改新密码
	insertPass := utils.BcryptHash(newPass)
	query2 := fmt.Sprintf("update %s set password  = ? where id = ?", m.table)

	_, err = m.conn.Exec(query2, insertPass, userId)
	if err != nil {
		return err
	}
	return nil

}
