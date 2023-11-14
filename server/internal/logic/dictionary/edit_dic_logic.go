package dictionary

import (
	"context"
	"zero-server/server/model/admin_dic"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditDicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditDicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditDicLogic {
	return &EditDicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditDicLogic) EditDic(req *types.EditDicReq) (resp *types.EditDicResp, err error) {
	resp = new(types.EditDicResp)
	DicTionary := &admin_dic.SysDictionaries{
		Id:     int64(req.Id),
		Name:   req.Name,
		Type:   req.Type,
		Status: req.Status,
		Desc:   req.Desc,
	}
	err = l.svcCtx.DicModel.Update(l.ctx, DicTionary)
	if err != nil {
		return nil, err
	}
	resp.Msg = "操作成功"

	return resp, nil
}
