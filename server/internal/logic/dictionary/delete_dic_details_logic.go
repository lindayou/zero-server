package dictionary

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDicDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDicDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDicDetailsLogic {
	return &DeleteDicDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDicDetailsLogic) DeleteDicDetails(req *types.DeleteDicDetailsReq) (resp *types.DeleteDicDetailsResp, err error) {
	resp = new(types.DeleteDicDetailsResp)
	err = l.svcCtx.DicDetail.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp.Msg = "删除成功"

	return resp, nil
}
