package menu

import (
	"context"
	menu2 "zero-server/server/model/menu"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (resp *types.DeleteMenuResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeleteMenuResp)
	menu := menu2.SysBaseMenus{
		Id: req.Id,
	}
	err = l.svcCtx.MenuModel.Delete(l.ctx, menu.Id)
	if err != nil {
		return nil, err
	}
	resp.Message = "删除成功"
	return resp, nil
}
