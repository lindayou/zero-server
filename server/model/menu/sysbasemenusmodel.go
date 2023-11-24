package menu

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var _ SysBaseMenusModel = (*customSysBaseMenusModel)(nil)
var (
	sysBaseMenusRowsWithPlaceHoldered = strings.Join(stringx.Remove(sysBaseMenusFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`", "`menu_level`", "`component`", "`active_name`", "`default_menu`", "`parent_id`"), "=?,") + "=?"
)

type (
	// SysBaseMenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysBaseMenusModel.
	SysBaseMenusModel interface {
		sysBaseMenusModel
		FindMenuList(ctx context.Context) ([]*SysBaseMenus, error)
		UpdateMenu(ctx context.Context, data *SysBaseMenus) error
		GetAuthorityMenu(ctx context.Context, AuthorityId int64) ([]*SysBaseMenus, error)
		AddMenuAuthority(ctx context.Context, menuIds []int, AuthorityId int64) error
		GetUserMenus(ctx context.Context, UserId int) ([]*SysBaseMenus, error)
	}

	customSysBaseMenusModel struct {
		*defaultSysBaseMenusModel
	}
)

// NewSysBaseMenusModel returns a model for the database table.
func NewSysBaseMenusModel(conn sqlx.SqlConn) SysBaseMenusModel {
	return &customSysBaseMenusModel{
		defaultSysBaseMenusModel: newSysBaseMenusModel(conn),
	}
}

func (m *defaultSysBaseMenusModel) FindMenuList(ctx context.Context) ([]*SysBaseMenus, error) {
	var menus = []*SysBaseMenus{}
	query := fmt.Sprintf("select * from %s ", m.table)
	err := m.conn.QueryRowsCtx(ctx, &menus, query)
	if err != nil {
	}
	return menus, nil

}
func (m *defaultSysBaseMenusModel) UpdateMenu(ctx context.Context, data *SysBaseMenus) error {
	query := fmt.Sprintf("update %s set %s where id =? ", m.table, sysBaseMenusRowsWithPlaceHoldered)
	_, err := m.conn.ExecCtx(ctx, query, data.Path, data.Name, data.Hidden, data.Sort, data.KeepAlive, data.Title, data.Icon, data.CloseTab, data.Id)
	if err != nil {
		return err
	}
	return err

}

func (m *defaultSysBaseMenusModel) GetAuthorityMenu(ctx context.Context, AuthorityId int64) ([]*SysBaseMenus, error) {
	var menus = []*SysBaseMenus{}
	query := fmt.Sprintf("select * FROM sys_base_menus where id in(select sys_base_menu_id from sys_authority_menus where sys_authority_authority_id = ?)")
	err := m.conn.QueryRowsCtx(ctx, &menus, query, AuthorityId)
	if err != nil {
		return nil, err
	}
	return menus, nil

}

func (m *defaultSysBaseMenusModel) AddMenuAuthority(ctx context.Context, menuIds []int, AuthorityId int64) error {
	err := m.conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		query := fmt.Sprintf("DELETE  from sys_authority_menus  WHERE  sys_authority_authority_id =? ")
		_, err := m.conn.ExecCtx(ctx, query, AuthorityId)
		if err != nil {
			return err
		}
		//批量插入
		query = fmt.Sprintf("insert into sys_authority_menus (sys_base_menu_id,sys_authority_authority_id) VALUES(?,?)")
		blk, err := sqlx.NewBulkInserter(m.conn, query)
		if err != nil {
			panic(err)
		}
		defer blk.Flush()

		for _, menuId := range menuIds {
			err := blk.Insert(menuId, AuthorityId)
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

func (m *defaultSysBaseMenusModel) GetUserMenus(ctx context.Context, UserId int) ([]*SysBaseMenus, error) {
	var menus = []*SysBaseMenus{}
	query := fmt.Sprintf("SELECT * from sys_base_menus WHERE id in (SELECT  DISTINCT(sys_base_menu_id ) from sys_authority_menus WHERE  sys_authority_authority_id in (  SELECT sys_authority_authority_id from sys_user_authority WHERE sys_user_id =?))")
	err := m.conn.QueryRowsCtx(ctx, &menus, query, UserId)
	if err != nil {
		return nil, err
	}
	return menus, nil

}
