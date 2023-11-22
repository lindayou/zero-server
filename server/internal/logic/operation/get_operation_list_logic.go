package operation

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOperationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOperationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationListLogic {
	return &GetOperationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOperationListLogic) GetOperationList(req *types.GetOperationListReq) (resp *types.GetOperationListResp, err error) {
	resp = new(types.GetOperationListResp)
	if req.Page < 1 {
		req.Page = 1
	}
	offset := (req.Page - 1) * req.PageSize

	list, total, err := l.svcCtx.Operation.GetOperationList(l.ctx, req.PageSize, offset)
	if err != nil {
		return nil, err
	}
	operations := make([]*types.Operation, 0)
	for _, item := range list {
		operations = append(operations, &types.Operation{
			Id:           item.Id,
			CreatedAt:    item.CreatedAt.Unix(),
			UpdatedAt:    item.UpdatedAt.Unix(),
			Ip:           item.Ip,
			Method:       item.Method,
			Path:         item.Path,
			Status:       item.Status,
			Latency:      item.Latency,
			Agent:        item.Agent,
			ErrorMessage: item.ErrorMessage,
			Body:         item.Body,
			Resp:         item.Resp,
			UserId:       item.UserId,
		})

	}
	resp.OperationList = operations
	resp.Total = total

	return resp, nil
}
