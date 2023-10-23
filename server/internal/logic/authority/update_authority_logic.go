package authority

import (
	"context"
	"zero-server/server/model/anthority_model"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAuthorityLogic {
	return &UpdateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAuthorityLogic) UpdateAuthority(req *types.UpdateAuthorityReq) (resp *types.UpdateAuthorityResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdateAuthorityResp)
	authority := anthority_model.SysAuthorities{
		AuthorityId:   req.AuthorityId,
		AuthorityName: req.AuthorityName,
		ParentId:      req.ParentId,
	}
	err = l.svcCtx.Authority.Update(l.ctx, &authority)
	if err != nil {
		return nil, err
	}
	resp.Message = "修改成功"
	return resp, nil
}
