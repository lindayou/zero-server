package user

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserAuthorityLogic {
	return &SetUserAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserAuthorityLogic) SetUserAuthority(req *types.SetUserAuthorityReq) (resp *types.SetUserAuthorityResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.SetUserAuthorityResp)
	err = l.svcCtx.AuthUser.SetUserAuth(l.ctx, req.AuthorityId, req.UserId)
	if err != nil {
		return nil, err
	}
	resp.Message = "设置成功"
	return resp, nil
}
