package dictionary

import (
	"context"
	"zero-server/server/model/admin_dic"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDicLogic {
	return &AddDicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDicLogic) AddDic(req *types.AddDicReq) (resp *types.AddDicResp, err error) {
	DicTionary := &admin_dic.SysDictionaries{
		Name:   req.Name,
		Type:   req.Type,
		Status: req.Status,
		Desc:   req.Desc,
	}
	insert, err := l.svcCtx.DicModel.Insert(l.ctx, DicTionary)
	if err != nil {
		return nil, err
	}
	resp = new(types.AddDicResp)
	resp.Msg = "插入成功"
	resp.Id, err = insert.LastInsertId()

	return resp, nil
}
