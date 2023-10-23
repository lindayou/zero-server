package menu

import (
	"context"
	menu2 "zero-server/server/model/menu"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditMenuLogic {
	return &EditMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditMenuLogic) EditMenu(req *types.EditMenuReq) (resp *types.EditMenuResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.EditMenuResp)
	menu := menu2.SysBaseMenus{
		Id:        req.Id,
		Path:      req.Path,
		Name:      req.Name,
		Hidden:    int64(req.Hidden),
		Sort:      int64(req.Sort),
		KeepAlive: int64(req.KeepAlive),
		Title:     req.Title,
		Icon:      req.Icon,
		CloseTab:  int64(req.CloseTab),
	}
	err = l.svcCtx.MenuModel.UpdateMenu(l.ctx, &menu)
	if err != nil {
		return nil, err
	}
	resp.Message = "修改成功"
	return resp, nil
}
