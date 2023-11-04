package menu

import (
	"context"
	"strconv"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(req *types.GetMenuAuthorityReq) (resp *types.GetMenuAuthorityResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetMenuAuthorityResp)
	//根据角色Id 获取菜单ID 再获取菜单
	menus, err := l.svcCtx.MenuModel.GetAuthorityMenu(l.ctx, req.AuthorityId)
	//获取全部菜单后进行父子返回
	treeMap := make(map[string][]*types.Menu)
	for _, menu := range menus {
		treeMap[menu.ParentId] = append(treeMap[menu.ParentId], &types.Menu{
			Id:        int(menu.Id),
			CreatedAt: menu.CreatedAt.Unix(),
			UpdatedAt: menu.UpdatedAt.Unix(),
			ParentId:  menu.ParentId,
			Path:      menu.Path,
			Name:      menu.Name,
			Hidden:    menu.Hidden,
			Sort:      int(menu.Sort),
			KeepAlive: menu.KeepAlive,
			Title:     menu.Title,
			Icon:      menu.Icon,
			CloseTab:  int(menu.CloseTab),
			Children:  nil,
		})
	}

	respList := make([]*types.Menu, 0)
	respList = append(respList, treeMap["0"]...)

	for _, item := range respList {
		getTreeList(item, treeMap)
	}
	resp.MenuList = respList

	return resp, nil
}

func getTreeList(menu *types.Menu, treeMap map[string][]*types.Menu) {
	menu.Children = treeMap[strconv.Itoa(menu.Id)]
	for _, item := range menu.Children {
		getTreeList(item, treeMap)
	}
}
