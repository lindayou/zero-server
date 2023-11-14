package dictionary

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDicDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDicDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDicDetailsLogic {
	return &GetDicDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDicDetailsLogic) GetDicDetails(req *types.GetDicDetailsReq) (resp *types.GetDicDetailsResp, err error) {
	resp = new(types.GetDicDetailsResp)
	DetailsList := make([]*types.DicDetail, 0)
	details, err := l.svcCtx.DicModel.GetDicDetails(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	for _, details := range details {
		DetailsList = append(DetailsList, &types.DicDetail{
			Id:              details.Id,
			SysDictionatyId: details.SysDictionaryId,
			Date:            int(details.CreatedAt.Unix()),
			Lable:           details.Label,
			Value:           details.Value,
			Status:          details.Status,
			Sort:            details.Sort,
		})

	}
	resp.DetailsList = DetailsList

	return resp, nil
}
