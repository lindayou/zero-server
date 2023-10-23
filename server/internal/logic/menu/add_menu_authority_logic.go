package menu

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuAuthorityLogic {
	return &AddMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuAuthorityLogic) AddMenuAuthority(req *types.AddMenuAuthorityReq) (resp *types.AddMenuAuthorityResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.AddMenuAuthorityResp)
	ids := []int{}
	for _, item := range req.MenuList {
		ids = append(ids, item.Id)
	}
	err = l.svcCtx.MenuModel.AddMenuAuthority(l.ctx, ids, req.AuthorityId)
	if err != nil {
		return nil, err
	}
	resp.Message = "设置成功"
	return resp, nil
}
