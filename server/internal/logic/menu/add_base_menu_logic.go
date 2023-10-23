package menu

import (
	"context"
	"zero-server/server/model/menu"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBaseMenuLogic {
	return &AddBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBaseMenuLogic) AddBaseMenu(req *types.AddMenuReq) (resp *types.AddMenuResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.AddMenuResp)
	baseMenu := menu.SysBaseMenus{
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Hidden:    int64(req.Hidden),
		Sort:      int64(req.Sort),
		KeepAlive: int64(req.KeepAlive),
		Title:     req.Title,
		Icon:      req.Icon,
		CloseTab:  int64(req.CloseTab),
	}
	_, err = l.svcCtx.MenuModel.Insert(l.ctx, &baseMenu)
	if err != nil {
		return nil, err
	}
	resp.Message = "添加成功"
	return resp, nil
}