// Code generated by goctl. DO NOT EDIT.

package anthority_model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysAuthoritiesFieldNames          = builder.RawFieldNames(&SysAuthorities{})
	sysAuthoritiesRows                = strings.Join(sysAuthoritiesFieldNames, ",")
	sysAuthoritiesRowsExpectAutoSet   = strings.Join(stringx.Remove(sysAuthoritiesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysAuthoritiesRowsWithPlaceHolder = strings.Join(stringx.Remove(sysAuthoritiesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sysAuthoritiesModel interface {
		Insert(ctx context.Context, data *SysAuthorities) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysAuthorities, error)
		Update(ctx context.Context, data *SysAuthorities) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysAuthoritiesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysAuthorities struct {
		Id            int64        `db:"id"`
		CreatedAt     sql.NullTime `db:"created_at"`
		UpdatedAt     sql.NullTime `db:"updated_at"`
		DeletedAt     sql.NullTime `db:"deleted_at"`
		AuthorityId   int64        `db:"authority_id"`
		AuthorityName string       `db:"authority_name"` // 角色名
		ParentId      int64        `db:"parent_id"`      // 父角色ID
		DefaultRouter string       `db:"default_router"` // 默认菜单
	}
)

func newSysAuthoritiesModel(conn sqlx.SqlConn) *defaultSysAuthoritiesModel {
	return &defaultSysAuthoritiesModel{
		conn:  conn,
		table: "`sys_authorities`",
	}
}

func (m *defaultSysAuthoritiesModel) withSession(session sqlx.Session) *defaultSysAuthoritiesModel {
	return &defaultSysAuthoritiesModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`sys_authorities`",
	}
}

func (m *defaultSysAuthoritiesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysAuthoritiesModel) FindOne(ctx context.Context, id int64) (*SysAuthorities, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysAuthoritiesRows, m.table)
	var resp SysAuthorities
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysAuthoritiesModel) Insert(ctx context.Context, data *SysAuthorities) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, sysAuthoritiesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.AuthorityId, data.AuthorityName, data.ParentId, data.DefaultRouter)
	return ret, err
}

func (m *defaultSysAuthoritiesModel) Update(ctx context.Context, data *SysAuthorities) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysAuthoritiesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.AuthorityId, data.AuthorityName, data.ParentId, data.DefaultRouter, data.Id)
	return err
}

func (m *defaultSysAuthoritiesModel) tableName() string {
	return m.table
}
