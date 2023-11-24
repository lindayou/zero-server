package menu

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenusLogic {
	return &GetUserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenusLogic) GetUserMenus(req *types.GetUserMenusReq) (resp *types.GetUserMenusResp, err error) {
	resp = new(types.GetUserMenusResp)
	userId := req.UserId
	menus, err := l.svcCtx.MenuModel.GetUserMenus(l.ctx, userId)
	if err != nil {
		return nil, err
	}
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
