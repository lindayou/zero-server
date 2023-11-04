package menu

import (
	"context"
	"strconv"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList(req *types.GetMenuListReq) (resp *types.GetMenuListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetMenuListResp)
	dbMenuLists, err := l.svcCtx.MenuModel.FindMenuList(l.ctx)
	if err != nil {
		return nil, err
	}
	var menuLists = []*types.Menu{}

	// parentId 为key  其对应的menuList为value
	totalMap := make(map[string][]*types.Menu)
	for _, menuList := range dbMenuLists {
		totalMap[menuList.ParentId] = append(totalMap[menuList.ParentId], &types.Menu{
			Id:        int(menuList.Id),
			CreatedAt: menuList.CreatedAt.Unix(),
			UpdatedAt: menuList.UpdatedAt.Unix(),
			ParentId:  menuList.ParentId,
			Path:      menuList.Path,
			Name:      menuList.Name,
			Hidden:    menuList.Hidden,
			Sort:      int(menuList.Sort),
			KeepAlive: menuList.KeepAlive,
			Icon:      menuList.Icon,
			CloseTab:  int(menuList.CloseTab),
			Title:     menuList.Title,
		})
	}
	//先拿到parentId为0的List  这组为第一层菜单
	menuLists = totalMap["0"]
	for _, firstMenu := range menuLists {
		//插入子菜单
		getBaseChildrenList(firstMenu, totalMap)
	}
	resp.MenuList = menuLists
	//遍历这一组并拿到对应的ID 找到父ID为这一组id的分别放入其子菜单
	return resp, nil
}

func getBaseChildrenList(menu2 *types.Menu, totalMap map[string][]*types.Menu) {
	menu2.Children = totalMap[strconv.Itoa(menu2.Id)]
	for _, child := range menu2.Children {
		getBaseChildrenList(child, totalMap)
	}
}
