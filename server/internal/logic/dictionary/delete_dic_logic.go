package dictionary

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDicLogic {
	return &DeleteDicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDicLogic) DeleteDic(req *types.DeleteDicReq) (resp *types.DeleteDicResp, err error) {
	resp = new(types.DeleteDicResp)
	err = l.svcCtx.DicModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp.Msg = "删除成功"
	return resp, nil
}
