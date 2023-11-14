package dictionary

import (
	"context"
	"zero-server/server/model/admin_dic_detail"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditDicDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditDicDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditDicDetailsLogic {
	return &EditDicDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditDicDetailsLogic) EditDicDetails(req *types.EditDicDetailsReq) (resp *types.EditDicDetailsResp, err error) {
	resp = new(types.EditDicDetailsResp)
	EditDicDetails := &admin_dic_detail.SysDictionaryDetails{
		Id:              req.Id,
		Label:           req.Lable,
		Value:           req.Value,
		Status:          req.Status,
		Sort:            req.Sort,
		SysDictionaryId: req.SysDictionatyId,
	}
	err = l.svcCtx.DicDetail.Update(l.ctx, EditDicDetails)
	if err != nil {
		return nil, err
	}
	resp.Msg = "修改成功"
	return resp, nil
}
