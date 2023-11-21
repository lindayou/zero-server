package operation

import (
	"context"
	"zero-server/server/model/admin_operation"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOperationLogic {
	return &CreateOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOperationLogic) CreateOperation(req *types.CreateOperationReq) (resp *types.CreateOperationResp, err error) {
	resp = new(types.CreateOperationResp)
	insertOperation := &admin_operation.SysOperationRecords{
		Ip:           req.Ip,
		Method:       req.Method,
		Path:         req.Path,
		Status:       req.Status,
		Latency:      req.Latency,
		Agent:        req.Agent,
		ErrorMessage: req.ErrorMessage,
		Body:         req.Body,
		Resp:         req.Resp,
		UserId:       req.UserId,
	}
	_, err = l.svcCtx.Operation.Insert(l.ctx, insertOperation)
	if err != nil {
		return nil, err
	}
	resp.Msg = "操作成功"

	return resp, nil
}
