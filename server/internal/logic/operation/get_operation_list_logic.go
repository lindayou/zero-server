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
	// todo: add your logic here and delete this line

	return
}
