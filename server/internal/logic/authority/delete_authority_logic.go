package authority

import (
	"context"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAuthorityLogic {
	return &DeleteAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAuthorityLogic) DeleteAuthority(req *types.DeleteAuthorityReq) (resp *types.DeleteAuthorityResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeleteAuthorityResp)

	err = l.svcCtx.Authority.DeleteAuthority(l.ctx, req.AuthorityId)
	if err != nil {
		return nil, err
	}
	resp.Message = "删除成功"
	return resp, nil
}
